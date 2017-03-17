// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

// reactGen is a go generate generator that helps to automate the process of
// writing GopherJS React web applications.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/myitcv/gogenerate"
)

const (
	reactGenCmd = "reactGen"
)

var (
	fLicenseFile = gogenerate.LicenseFileFlag()
	fGoGenLog    = gogenerate.LogFlag()
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(reactGenCmd + ": ")

	defer func() {
		if err, ok := recover().(error); ok {
			log.Fatalln(err)
		}
	}()

	flag.Parse()

	gogenerate.DefaultLogLevel(fGoGenLog, gogenerate.LogFatal)

	envFileName, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	wd, err := os.Getwd()
	if err != nil {
		fatalf("unable to get working directory: %v", err)
	}

	// are we running against the first file that contains the reactGen directive?
	// if not return
	dirFiles, err := gogenerate.FilesContainingCmd(wd, reactGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if len(dirFiles) == 0 {
		fatalf("cannot find any files containing the %v directive", reactGenCmd)
	}

	if envFileName != dirFiles[0] {
		return
	}

	license, err := gogenerate.CommentLicenseHeader(fLicenseFile)
	if err != nil {
		fatalf("could not comment license file: %v", err)
	}

	// if we get here, we know we are the first file...

	dogen(wd, license)
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func infof(format string, args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Printf(format, args...)
	}
}
