/*

stateGen is a go generate generator that helps to automate the process of
creating state trees for use in GopherJS React web applications.

For more information see https://github.com/myitcv/react/wiki

*/
package main // import "myitcv.io/react/cmd/stateGen"

import (
	"flag"
	"fmt"
	"log"
	"os"

	"myitcv.io/gogenerate"
)

const (
	stateGenCmd = "stateGen"
)

var (
	fLicenseFile = gogenerate.LicenseFileFlag()
	fGoGenLog    = gogenerate.LogFlag()
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(stateGenCmd + ": ")

	flag.Parse()

	gogenerate.DefaultLogLevel(fGoGenLog, gogenerate.LogFatal)

	envFile, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	wd, err := os.Getwd()
	if err != nil {
		fatalf("unable to get working directory: %v", err)
	}

	dirFiles, err := gogenerate.FilesContainingCmd(wd, stateGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if dirFiles == nil {
		fatalf("cannot find any files containing the %v directive", stateGenCmd)
	}

	if dirFiles[envFile] != 1 {
		fatalf("expected a single occurrence of %v directive in %v. Got: %v", stateGenCmd, envFile, dirFiles)
	}

	license, err := gogenerate.CommentLicenseHeader(fLicenseFile)
	if err != nil {
		fatalf("could not comment license file: %v", err)
	}

	// if we get here, we know we are the first file...

	dogen(os.Stderr, wd, license)
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func infof(format string, args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Printf(format, args...)
	}
}
