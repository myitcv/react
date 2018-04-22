package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ConfigFileName = ".ggconfig.json"
)

type Config struct {
	Typed   []string
	Untyped []string

	// maps of the packages
	typed   map[string]struct{}
	untyped map[string]struct{}

	// maps of the commands, essentially the bases of the packages
	typedCmds   map[string]struct{}
	untypedCmds map[string]struct{}
}

var config Config

var validCmd = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func loadConfig() {
	if *fUntyped != "" || *fTyped != "" {
		config.Untyped = splitCmdList(*fUntyped)
		config.Typed = splitCmdList(*fTyped)
	} else {
		// TODO maybe instead of using $PWD as the starting point for finding a config file we should start at
		// the package directory...

		var err error

		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		dir := wd

		var fi *os.File

		for {
			f := filepath.Join(dir, ConfigFileName)

			fi, err = os.Open(f)
			if err == nil {
				break
			}

			p := filepath.Dir(dir)

			if p == dir {
				break
			}

			dir = p
		}

		if fi == nil {
			log.Fatalf("Could not find %v in %v (or any parent directory)", ConfigFileName, wd)
		}

		j := json.NewDecoder(fi)
		err = j.Decode(&config)
		if err != nil {
			log.Fatalf("Could not decode config file %v:\n%v", fi.Name(), err)
		}
	}

	config.typed = make(map[string]struct{})
	config.untyped = make(map[string]struct{})

	config.typedCmds = make(map[string]struct{})
	config.untypedCmds = make(map[string]struct{})

	for _, v := range config.Typed {
		b := filepath.Base(v)
		config.typed[v] = struct{}{}
		config.typedCmds[b] = struct{}{}
	}

	for _, v := range config.Untyped {
		b := filepath.Base(v)
		config.untyped[v] = struct{}{}
		config.untypedCmds[b] = struct{}{}
	}

	config.Typed = keySlice(config.typed)
	config.Untyped = keySlice(config.untyped)
}

func splitCmdList(s string) []string {
	s = strings.TrimSpace(s)
	ps := strings.Split(s, ",")

	parts := make([]string, 0, len(ps))

	for _, v := range ps {
		v = strings.TrimSpace(v)

		if v == "" {
			continue
		}

		if !validCmd.MatchString(v) {
			log.Fatalf("Invalid go generate cmd: %v\n", v)
		}

		parts = append(parts, v)
	}

	return parts
}
