// +build !js

package main

import (
	"testing"
)

func TestSuccessNoOutput(t *testing.T) {
	r := testRunner(t, "test.001")
	r.run()
	r.exitCode(0)
	r.grepBoth("ok", "failed to find test pass message")
}

func TestSuccessOutput(t *testing.T) {
	r := testRunner(t, "test.002")
	r.run()
	r.exitCode(0)
	r.grepBoth("Some output", "failed to find output")
}

func TestFail(t *testing.T) {
	r := testRunner(t, "test.003")
	r.run()
	r.exitCode(1)
	r.grepBoth("failed for no reason", "failed to find fail output")
}
