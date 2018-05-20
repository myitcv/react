// gjbt is a simple (temporary) wrapper for GopherJS to run tests in Chrome as
// opposed to NodeJS.
package main // import "myitcv.io/gjbt"

import (
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/kisielk/gotool"
	"github.com/sclevine/agouti"
)

type res struct {
	Error    string
	ExitCode int
}

var (
	fTags   = flag.String("tags", "", "tags to pass to the GopherJS compiler")
	fBinary = flag.String("binary", "", "path to Chrome binary")
)

// TODO only works for Chrome for now

// TODO support verbose mode in some way

func main() {
	flag.Parse()

	// for each package:
	//
	// 1. gopherjs test -c -o /tmp/file
	// 2. Run the below and

	opts := []agouti.Option{
		agouti.ChromeOptions(
			"args", []string{
				"headless",
				"no-default-browser-check",
				"verbose",
				"no-sandbox",
				"no-first-run",
				"disable-default-apps",
				"disable-popup-blocking",
				"disable-translate",
				"disable-background-timer-throttling",
				"disable-renderer-backgrounding",
				"disable-device-discovery-notifications",
			},
		),
		agouti.Desired(
			agouti.Capabilities{
				"loggingPrefs": map[string]string{
					"browser": "INFO",
				},
			},
		),
	}

	if *fBinary != "" {
		opts = append(opts,
			agouti.ChromeOptions(
				"binary", *fBinary,
			))
	}

	driver := agouti.ChromeDriver(opts...)

	if err := driver.Start(); err != nil {
		panic(err)
	}

	pkgs := gotool.ImportPaths(flag.Args())

	failed := false

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		tf, err := ioutil.TempFile("", "")
		if err != nil {
			panic(err)
		}

		bpkg, err := build.Import(pkg, wd, build.FindOnly)
		if err != nil {
			panic(err)
		}

		args := []string{"test", "--tags", *fTags, "-c", "-o", tf.Name()}

		args = append(args, pkg)

		cmd := exec.Command("gopherjs", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			os.Remove(tf.Name())
			fmt.Printf("%v\n", err)
			failed = true
			continue
		}

		test, err := ioutil.ReadFile(tf.Name())
		if err != nil {
			panic(err)
		}

		os.Remove(tf.Name())

		p, err := driver.NewPage()
		if err != nil {
			panic(err)
		}

		var ec res

		status := "ok  "
		start := time.Now()

		err = p.RunScript(`try {
			`+string(test)+`
		}
		catch (e) {
			window.$GopherJSTestResult = {
				Error: e.toString(),
				ExitCode: 1,
			};
		};
		if(typeof window.$GopherJSTestResult === 'number') {
			window.$GopherJSTestResult = {
				ExitCode: window.$GopherJSTestResult
			}
		};
		return window.$GopherJSTestResult`, nil, &ec)

		if err != nil {
			panic(err)
		}

		if ec.ExitCode != 0 {
			status = "FAIL"
			failed = true
		}

		if ec.Error != "" {
			fmt.Println(ec.Error)
		}
		fmt.Printf("%s\t%s\t%.3fs\n", status, bpkg.ImportPath, time.Since(start).Seconds())

		logs, err := p.ReadNewLogs("browser")
		if err != nil {
			panic(err)
		}

		for _, l := range logs {
			parts := strings.SplitN(l.Message, " ", 3)

			line := parts[2]

			if strings.HasPrefix(line, "\"") && strings.HasSuffix(line, "\"") {
				l, err := strconv.Unquote(parts[2])
				if err != nil {
					panic(err)
				}
				line = l
			}

			fmt.Println(line)
		}
	}

	if err := driver.Stop(); err != nil {
		panic(err)
	}

	if failed {
		os.Exit(1)
	}
}
