// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main // import "myitcv.io/immutable/cmd/immutableGen"

import (
	"flag"
	"log"
	"os"

	"myitcv.io/gogenerate"
)

const (
	immutableGenCmd = "immutableGen"
)

var (
	fGoGenCmds   gogenCmds
	fLicenseFile = gogenerate.LicenseFileFlag()
	fGoGenLog    = gogenerate.LogFlag()
)

func init() {
	flag.Var(&fGoGenCmds, "G", "Path to search for imports (flag can be used multiple times)")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix(immutableGenCmd + ": ")

	gogenerate.DefaultLogLevel(fGoGenLog, gogenerate.LogFatal)

	envFile, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	envPkgName, ok := os.LookupEnv(gogenerate.GOPACKAGE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOPACKAGE)
	}

	wd, err := os.Getwd()
	if err != nil {
		fatalf("unable to get working directory: %v", err)
	}

	dirFiles, err := gogenerate.FilesContainingCmd(wd, immutableGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if len(dirFiles) == 0 {
		fatalf("cannot find any files containing the %v directive", immutableGenCmd)
	}

	if envFile != dirFiles[0] {
		return
	}

	licenseHeader, err := gogenerate.CommentLicenseHeader(fLicenseFile)
	if err != nil {
		fatalf("could not comment license file: %v", err)
	}

	execute(wd, envPkgName, licenseHeader, fGoGenCmds)
}
