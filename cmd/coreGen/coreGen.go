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

	envFileName, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	envPkg, ok := os.LookupEnv(gogenerate.GOPACKAGE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOPACKAGE)
	}

	// are we running against the first file that contains the coreGen directive?
	// if not return
	dirFiles, err := gogenerate.FilesContainingCmd(wd, coreGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if len(dirFiles) == 0 {
		fatalf("cannot find any files containing the %v directive", coreGenCmd)
	}

	if envFileName != dirFiles[0] {
		return
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
