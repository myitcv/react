// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

/*

coreGen is a go generate generator that helps to automate writing the core of
myitcv.io/react.

For more information see https://github.com/myitcv/react/wiki

*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"myitcv.io/gogenerate"
)

const (
	coreGenCmd = "coreGen"

	jsPkg = "github.com/gopherjs/gopherjs/js"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(coreGenCmd + ": ")

	flag.Usage = usage
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		fatalf("unable to get working directory: %v", err)
	}

	mainGen(wd)
}

func mainGen(wd string) {
	gogenerate.DefaultLogLevel(fGoGenLog, gogenerate.LogFatal)

	envFile, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	envPkg, ok := os.LookupEnv(gogenerate.GOPACKAGE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOPACKAGE)
	}

	dirFiles, err := gogenerate.FilesContainingCmd(wd, coreGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if dirFiles == nil {
		fatalf("cannot find any files containing the %v directive", coreGenCmd)
	}

	if dirFiles[envFile] != 1 {
		fatalf("expected a single occurrence of %v directive in %v. Got: %v", coreGenCmd, envFile, dirFiles)
	}

	license, err := gogenerate.CommentLicenseHeader(fLicenseFile)
	if err != nil {
		fatalf("could not comment license file: %v", err)
	}

	// if we get here, we know we are the first file...

	dogen(wd, envPkg, license)
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func infof(format string, args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Printf(format, args...)
	}
}
