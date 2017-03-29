package main

import (
	"flag"
	"fmt"
	"os"
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

var fInit initFlag

func init() {
	flag.Var(&fInit, "init", "create a GopherJS React application using the specified template (see below)")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\t%v [-init <template>] [-gglog <log_level>] [-licenseFile <filepath>]\n\n", os.Args[0])

	flag.PrintDefaults()

	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "The flag -init only understands a single value for now: minimal. This is a minimal")
	fmt.Fprintln(os.Stderr, "Gopher React application.")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "When -init is not specified, it is assumed that reactGen is being called indirectly")
	fmt.Fprintln(os.Stderr, "via go generate. The options for -gglog and -licenseFile would therefore be set in")
	fmt.Fprintln(os.Stderr, "via the //go:generate directives. See https://blog.golang.org/generate for more details.")
}
