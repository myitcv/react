package main

import (
	"flag"
	"fmt"
	"os"

	"myitcv.io/gogenerate"
)

var (
	fLicenseFile = gogenerate.LicenseFileFlag()
	fGoGenLog    = gogenerate.LogFlag()
	fCore        = flag.Bool("core", false, "indicates we are generating for a core component (only do props expansion)")
	fInit        initFlag
)

type initFlag struct {
	val *string
}

func (f *initFlag) String() string {
	return "(does not have a default value)"
}

func (f *initFlag) Set(s string) error {
	f.val = &s
	return nil
}

func init() {
	flag.Var(&fInit, "init", "create a GopherJS React application using the specified template (see below)")
}

func usage() {
	f := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, format, args...)
	}

	l := func(args ...interface{}) {
		fmt.Fprintln(os.Stderr, args...)
	}

	l("Usage:")
	f("\t%v [-init <template>]\n", os.Args[0])
	f("\t%v [-gglog <log_level>] [-licenseFile <filepath>] [-core]\n", os.Args[0])
	l()

	flag.PrintDefaults()
}
