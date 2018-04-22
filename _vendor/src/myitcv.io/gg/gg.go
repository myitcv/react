// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main // import "myitcv.io/gg"

// gg is a wrapper for ``go generate''. More docs to follow

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/kisielk/gotool"
	"myitcv.io/gogenerate"
)

const (
	untypedLoopLimit = 10
	typedLoopLimit   = untypedLoopLimit
)

var (
	wd string
)

// All code basically derived from rsc.io/gt

// TODO we effectively read from some files twice... whilst computing stale and scanning
// for directives. These two operations could potentially be collapsed into a single read

func main() {
	var err error

	log.SetFlags(0)
	log.SetPrefix("gg: ")

	defer func() {
		err := recover()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	flag.Parse()

	wd, err = os.Getwd()
	if err != nil {
		fatalf("could not get working directory: %v", err)
	}

	loadConfig()

	specs := gotool.ImportPaths(flag.Args())
	sort.Strings(specs)

	readPkgs(specs, true)

	pkgs := make([]string, 0, len(pkgInfo))
	for k := range pkgInfo {
		pkgs = append(pkgs, k)
	}

	pkgs = cmdList(pkgs)

	if len(pkgs) == 0 {
		vvlogf("No packages contain any directives")
		os.Exit(0)
	}

	if *fList {
		// cmdList above will have done the logging for us

		os.Exit(0)
	}

	untypedRunExp := buildGoGenRegex(config.Untyped)
	typedRunExp := buildGoGenRegex(config.Typed)

	diffs := computeStale(pkgs, false)

	typedCount := 1

	for {
		untypedCount := 1

		preUntyped := snapHash(diffs)

		for len(diffs) > 0 {
			if untypedCount > untypedLoopLimit {
				fatalf("Exceeded loop limit for untyped go generate cmd: %v\n", untypedRunExp)
			}

			vvlogf("Untyped iteration %v.%v\n", typedCount, untypedCount)
			goGenerate(diffs, untypedRunExp)
			untypedCount++

			// order is significant here... because the computeStale
			// call does a readPkgs
			prevDiffs := diffs
			diffs = computeStale(prevDiffs, true)
			cmdList(prevDiffs)
		}

		// TODO work out what to do here when gg is being used in conjunction
		// with gai
		t := time.Now()
		vvlogf("pre go install")
		suc, _ := goInstall(pkgs)
		vvlogf("post go install %v", time.Now().Sub(t))

		if len(suc) == 0 {
			fatalf("No packages from %v succeeded install; cannot continue\n", pkgs)
		}

		if typedCount > typedLoopLimit {
			fatalf("Exceeded loop limit for typed go generate cmd: %v\n", untypedRunExp)
		}

		vvlogf("Typed iteration %v.0\n", typedCount)
		goGenerate(suc, typedRunExp)
		typedCount++

		// order is significant here... because the computeStale
		// call does a readPkgs
		computeStale(suc, true)
		cmdList(suc)

		postTypedDelta := deltaHash(preUntyped)

		// if there has been no change then regardless of how many fails etc
		// we should break
		if len(postTypedDelta) == 0 {
			vvlogf("no delta from start of untyped iteration; breaking")
			break
		}

		pkgs = postTypedDelta
	}
}

func buildGoGenRegex(parts []string) string {
	escpd := make([]string, len(parts))

	for i := range parts {
		cmd := filepath.Base(parts[i])
		escpd[i] = regexp.QuoteMeta(cmd)
	}

	exp := fmt.Sprintf(gogenerate.GoGeneratePrefix+" (?:%v)(?:$| )", strings.Join(escpd, "|"))

	// aggressively ensure the regexp compiles here... else a call to go generate
	// will be useless
	_, err := regexp.Compile(exp)
	if err != nil {
		fatalf("Could not form valid go generate command: %v\n", err)
	}

	return exp
}

func goGenerate(pkgs []string, runExp string) {
	args := []string{"generate"}

	if *fVerbose {
		args = append(args, "-v")
	}

	if *fExecute {
		args = append(args, "-x")
	}

	args = append(args, "-run", runExp)
	args = append(args, pkgs...)

	xlogf("go %v", strings.Join(args, " "))

	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		fatalf("go generate: %v\n%s", err, out)
	}

	if len(out) > 0 {
		// we always log the output from go generate
		fmt.Print(string(out))
	}
}

func goInstall(pkgs []string) ([]string, []string) {
	fmap := make(map[string]struct{})

	xlogf("gai %v", strings.Join(pkgs, " "))
	vvlogf("gai %v", strings.Join(pkgs, " "))

	out, err := exec.Command("gai", pkgs...).CombinedOutput()
	if err != nil {
		sc := bufio.NewScanner(bytes.NewBuffer(out))
		for sc.Scan() {
			line := sc.Text()

			if strings.HasPrefix(line, "# ") {
				parts := strings.Fields(line)

				if len(parts) != 2 {
					fatalf("could not parse go install output\n%v", string(out))
				}

				fmap[parts[1]] = struct{}{}
			}
		}

		if err := sc.Err(); err != nil {
			fatalf("could not parse go install output\n%v", string(out))
		}
	}

	if len(out) > 0 {
		xlog(string(out))
	}

	var f, s []string

	for _, p := range pkgs {
		if _, ok := fmap[p]; ok {
			f = append(f, p)
		} else {
			s = append(s, p)
		}
	}

	return s, f
}

// cmdList returns a subset of packages (subset of pNames) that contain directives
// and a map[package] -> map[cmd]struct{} of which commands are used in which packages
// As it scans each package in pNames it removes any generated files that do not have
// an occurence of a directive for the associated generator in the package (not test
// aware right now). In the process it also validates the directives that are present
func cmdList(pNames []string) []string {
	cmds := make(map[string]map[string]struct{})

	for _, pName := range pNames {
		var h map[string]struct{}

		pkg := pkgInfo[pName]

		var goFiles []string
		goFiles = append(goFiles, pkg.GoFiles...)
		goFiles = append(goFiles, pkg.CgoFiles...)
		goFiles = append(goFiles, pkg.TestGoFiles...)
		goFiles = append(goFiles, pkg.XTestGoFiles...)

		cmdFiles := make(map[string][]string)

		for _, f := range goFiles {
			f = filepath.Join(pkg.Dir, f)

			visitDir := func(line int, dirArgs []string) error {
				if *fList {
					rel, err := filepath.Rel(wd, f)
					if err != nil {
						fatalf("could not create filepath.Re(%q, %q): %q", wd, f, err)
					}
					fmt.Printf("%v:%v: %v\n", rel, line, strings.Join(dirArgs, " "))
				}
				if h == nil {
					h = make(map[string]struct{})
					cmds[pName] = h
				}

				h[dirArgs[0]] = struct{}{}

				return nil
			}

			if cmd, ok := gogenerate.FileIsGenerated(f); ok {
				// we only care about cmds which we know about in our config
				// for now this helps to deal with the edge case that is protobuf
				// files

				_, oktyp := config.typedCmds[cmd]
				_, okuntyp := config.untypedCmds[cmd]

				if oktyp || okuntyp {
					cmdFiles[cmd] = append(cmdFiles[cmd], f)
				}
			}

			gogenerate.DirFunc(pName, f, visitDir)
		}

		removed := false

		for c, fs := range cmdFiles {
			if _, ok := h[c]; !ok {
				for _, f := range fs {
					vvlogf("removing %v", f)

					removed = true

					err := os.Remove(f)
					if err != nil {
						fatalf("could not remove %v: %v", f, err)
					}
				}
			}
		}

		if removed {
			readPkgs([]string{pName}, false)
		}
	}

	cm := cmdMap(cmds)

	for v := range cm {

		_, tok := config.typedCmds[v]
		_, uok := config.untypedCmds[v]

		if !tok && !uok {
			log.Fatalf("go generate directive command \"%v\" is not specified as either typed or untyped", v)
		}
	}

	dirPkgs := make([]string, 0, len(cmds))
	for k := range cmds {
		dirPkgs = append(dirPkgs, k)
	}

	return dirPkgs
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func xlog(args ...interface{}) {
	if *fVVerbose || *fExecute {
		log.Print(args...)
	}
}

func xlogf(format string, args ...interface{}) {
	if *fVVerbose || *fExecute {
		log.Print(args...)
	}
}

func vvlogf(format string, args ...interface{}) {
	if *fVVerbose {
		log.Printf(format, args...)
	}
}

func cmdMap(cmds map[string]map[string]struct{}) map[string]struct{} {
	allCmds := make(map[string]struct{})

	for _, m := range cmds {
		for k := range m {
			allCmds[k] = struct{}{}
		}
	}

	return allCmds
}

func keySlice(m map[string]struct{}) []string {
	res := make([]string, 0, len(m))

	for k := range m {
		res = append(res, k)
	}

	return res
}
