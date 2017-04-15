// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"testing"

	"myitcv.io/gogenerate"
	"myitcv.io/immutable/util"
)

const (
	TestFiles = "internal/coretest"
)

func TestBasic(t *testing.T) {
	license := "// My favourite license\n\n"
	echoCmd := `echo "hello world"` // need a command that will succeed with zero exit code

	tmpl := "core.go"

	execute(TestFiles, "coretest", license, gogenCmds{echoCmd})

	tmplFile := filepath.Join(TestFiles, tmpl)
	genFile, ok := gogenerate.NameFileFromFile(tmplFile, immutableGenCmd)

	if !ok {
		t.Fatalf("could not calculated generated name for %v", tmplFile)
	}

	genOut, err := os.Open(genFile)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, genOut)
	if err != nil {
		panic(err)
	}

	_, err = genOut.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, tmplFile, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		panic(err)
	}

	foundMyStruct := false
	foundMySlice := false
	foundMyMap := false

	for _, d := range f.Decls {

		gd, ok := d.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}

		for _, s := range gd.Specs {
			ts := s.(*ast.TypeSpec)

			name, ok := util.IsImmTmplAst(ts)

			if !ok {
				continue
			}

			switch name {
			case "MyStruct":
				foundMyStruct = true
			case "MySlice":
				foundMySlice = true
			case "MyMap":
				foundMyMap = true
			}

			if name == "MyStruct" {
				foundMyStruct = true
			}
		}
	}

	if !foundMyStruct {
		t.Errorf("did not find MyStruct in generated output")
	}
	if !foundMySlice {
		t.Errorf("did not find MySlice in generated output")
	}
	if !foundMyMap {
		t.Errorf("did not find myMap in generated output")
	}
}
