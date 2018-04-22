package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	fXPkgs    xPkgs
	fVVerbose = flag.Bool("vv", false, "output commands as they are executed")
	fList     = flag.Bool("l", false, "list go generate directive commands in packages")
	fVerbose  = flag.Bool("v", false, "print the names of packages and files as they are processed")
	fExecute  = flag.Bool("x", false, "print commands as they are executed")
	fUntyped  = flag.String("untyped", "", "a list of untyped generators to run")
	fTyped    = flag.String("typed", "", "a list of typed generators to run")
)

type xPkgs []string

func (i *xPkgs) Set(value string) error {
	if strings.HasPrefix(value, ".") {
		log.Fatal("Cannot use \".\" as an exclude pattern")
	}

	*i = append(*i, value)
	return nil
}

func (i *xPkgs) String() string {
	return fmt.Sprint(*i)
}

func init() {
	flag.Var(&fXPkgs, "X", "packages to exclude")
}
