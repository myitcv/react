// +build !js

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
	"testing"
)

var gjbt string

func TestMain(m *testing.M) {
	td, err := ioutil.TempDir("", "gjbtbin")
	if err != nil {
		failf("failed to create gjbt build dir: %v", err)
	}

	gjbt = filepath.Join(td, "gjbt")

	cmd := exec.Command("go", "build", "-o", gjbt, "myitcv.io/gjbt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		failf("failed to compile gjbt: %v\n%s", err, out)
	}
	defer func() {
		os.Remove(gjbt)
	}()

	m.Run()
}

func failf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

type testRunnerData struct {
	dir string
	t   *testing.T

	ran bool

	actExitCode int
	stdout      bytes.Buffer
	stderr      bytes.Buffer
}

func testRunner(t *testing.T, dir string) *testRunnerData {
	return &testRunnerData{
		dir: dir,
		t:   t,
	}
}

func (tr *testRunnerData) run(flagAndArgs ...string) {
	tr.t.Helper()

	if tr.ran {
		tr.t.Fatalf("tried to run run twice")
	}

	dir := filepath.Join(".", "testdata", tr.dir)

	args := []string{"-tags", "js"}

	if len(flagAndArgs) == 0 {
		args = append(args, ".")
	} else {
		args = append(args, flagAndArgs...)
	}

	cmd := exec.Command(gjbt, args...)
	cmd.Dir = filepath.Join(dir)
	cmd.Stdout = &tr.stdout
	cmd.Stderr = &tr.stderr

	var exitCode int

	err := cmd.Run()
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			tr.t.Fatalf("unexpected error: %v", err)
		}

		exitCode = ee.Sys().(syscall.WaitStatus).ExitStatus()
	}

	tr.actExitCode = exitCode
	tr.ran = true
}

func (tr *testRunnerData) exitCode(i int) {
	tr.t.Helper()

	if tr.actExitCode != i {
		tr.t.Fatalf("exit code; want %v; got %v", i, tr.actExitCode)
	}
}

func (tr *testRunnerData) doGrepMatch(match string, b *bytes.Buffer) bool {
	tr.t.Helper()
	if !tr.ran {
		tr.t.Fatal("testsuite error: grep called before run")
	}
	re := regexp.MustCompile(match)
	for _, ln := range bytes.Split(b.Bytes(), []byte{'\n'}) {
		if re.Match(ln) {
			return true
		}
	}
	return false
}

func (tr *testRunnerData) grepStderr(match, msg string) {
	tr.t.Helper()

	if !tr.doGrepMatch(match, &tr.stderr) {
		tr.t.Log(msg)
		tr.t.Logf("pattern %v not found in standard error", match)
		tr.t.FailNow()
	}
}

func (tr *testRunnerData) grepBoth(match, msg string) {
	tr.t.Helper()

	if !tr.doGrepMatch(match, &tr.stdout) && !tr.doGrepMatch(match, &tr.stderr) {
		tr.t.Log(msg)
		tr.t.Logf("pattern %v not found in standard output or standard error", match)
		tr.t.FailNow()
	}
}
