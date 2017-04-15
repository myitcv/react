package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"testing"

	"github.com/kisielk/gotool"
)

func TestImmutableVetter(t *testing.T) {

	var expected = `_testFiles/test.go:10:2: immutable struct field must be immutable type; []byte is not
_testFiles/test.go:12:2: immutable struct field must be immutable type; myitcv.io/immutable/cmd/immutableVet/_testFiles.MyIntf is not
_testFiles/test.go:13:2: immutable struct field must be immutable type; myitcv.io/immutable/cmd/immutableVet/_testFiles.MyType is not
_testFiles/test.go:20:6: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:24:7: template type myitcv.io/immutable/cmd/immutableVet/_testFiles._Imm_Dummy should never get used
_testFiles/test.go:25:5: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:27:7: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:28:7: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:29:7: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:31:14: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:31:23: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:32:9: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy found
_testFiles/test.go:32:9: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:35:9: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:35:18: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:36:9: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy found
_testFiles/test.go:36:9: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
_testFiles/test.go:39:9: construct using new() or generated constructors
_testFiles/test.go:39:10: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy found
_testFiles/test.go:46:8: should not be using _Name of *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy immutable type
_testFiles/test.go:50:8: Range() of immutable type must appear in a range statement or used with an ellipsis as the second argument to append
_testFiles/test.go:55:7: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.intS found
_testFiles/test.go:56:8: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.intS found
_testFiles/test.go:73:9: non-pointer value of immutable type *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy found
_testFiles/test.go:73:9: type should be *myitcv.io/immutable/cmd/immutableVet/_testFiles.Dummy
`

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	specs := gotool.ImportPaths([]string{
		"myitcv.io/immutable/cmd/immutableVet/_testFiles",
	})

	emsgs := vet(wd, specs)

	bfr := bytes.NewBuffer(nil)

	for _, msg := range emsgs {
		fmt.Fprintf(bfr, "%v\n", msg)
	}

	diff := strDiff(expected, bfr.String())
	if diff != "" {
		fmt.Println(bfr.String())
		t.Errorf("Expected no diff; got:\n%v", diff)
	}
}

func mustTmpFile(dir string, prefix string) *os.File {
	res, err := ioutil.TempFile(dir, prefix)

	if err != nil {
		panic(err)
	}

	return res
}

func strDiff(exp, act string) string {
	actFn := mustTmpFile("", "").Name()
	expFn := mustTmpFile("", "").Name()

	err := ioutil.WriteFile(actFn, []byte(act), 077)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(expFn, []byte(exp), 077)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("diff", "-wu", expFn, actFn)
	res, err := cmd.CombinedOutput()
	if err != nil {
		ec := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if ec.ExitStatus() != 1 {
			panic(err)
		}
	}

	return string(res)
}
