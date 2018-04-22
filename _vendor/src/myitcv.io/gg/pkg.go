package main

import (
	"crypto/sha1"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	pkgInfo = map[string]*Package{}
)

type Package struct {
	*build.Package

	pkgHash string
}

func readPkgs(pkgs []string, ignore bool) {

All:
	for _, pn := range pkgs {
		p, err := build.Import(pn, wd, 0)
		if err != nil {
			fatalf("could not load package %v: %v", pn, err)
		}

		if ignore {
			a := p.ImportPath

			for _, x := range fXPkgs {
				if strings.HasSuffix(x, "/...") {
					p := strings.TrimSuffix(x, "/...")

					if a == p || strings.HasPrefix(a, p+"/") {
						continue All
					}
				} else {
					if a == x {
						continue All
					}
				}

			}

		}

		pkgInfo[p.ImportPath] = &Package{Package: p}
	}
}

func snapHash(pkgs []string) map[string]string {
	prevHashes := make(map[string]string, len(pkgs))
	for _, p := range pkgs {
		v := ""

		if pkg, ok := pkgInfo[p]; ok {
			v = pkg.pkgHash
		}

		prevHashes[p] = v
	}

	return prevHashes
}

func computeStale(pkgs []string, read bool) []string {
	snap := snapHash(pkgs)

	if read {
		readPkgs(pkgs, false)
	}

	for _, pkg := range pkgs {
		computePkgHash(pkgInfo[pkg])
	}

	return deltaHash(snap)
}

func deltaHash(snap map[string]string) []string {
	var deltas []string

	for p := range snap {
		if snap[p] != pkgInfo[p].pkgHash {
			deltas = append(deltas, p)
		}
	}

	return deltas
}

func computePkgHash(p *Package) {
	h := sha1.New()

	fmt.Fprintf(h, "pkg %v\n", p.ImportPath)

	hashFiles(h, p.Dir, p.GoFiles)
	hashFiles(h, p.Dir, p.CgoFiles)
	hashFiles(h, p.Dir, p.CFiles)
	hashFiles(h, p.Dir, p.CXXFiles)
	hashFiles(h, p.Dir, p.MFiles)
	hashFiles(h, p.Dir, p.HFiles)
	hashFiles(h, p.Dir, p.SFiles)
	hashFiles(h, p.Dir, p.SwigFiles)
	hashFiles(h, p.Dir, p.SwigCXXFiles)
	hashFiles(h, p.Dir, p.SysoFiles)
	hashFiles(h, p.Dir, p.TestGoFiles)
	hashFiles(h, p.Dir, p.XTestGoFiles)

	hash := fmt.Sprintf("%x", h.Sum(nil))
	p.pkgHash = hash
}

func hashFiles(h io.Writer, dir string, files []string) {
	for _, file := range files {
		fn := filepath.Join(dir, file)
		f, err := os.Open(fn)
		if err != nil {
			fatalf("could not open file %v: %v\n", fn, err)
		}

		fmt.Fprintf(h, "file %s\n", file)
		n, _ := io.Copy(h, f)
		fmt.Fprintf(h, "%d bytes\n", n)

		f.Close()
	}
}
