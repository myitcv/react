package main

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"path/filepath"

	"github.com/neelance/astrewrite"
)

func main() {
	importPath := os.Args[1]

	pkg, err := build.Import(importPath, "", 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	files := make([]*ast.File, len(pkg.GoFiles))
	for i, name := range pkg.GoFiles {
		file, err := parser.ParseFile(fset, filepath.Join(pkg.Dir, name), nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		files[i] = file
	}

	typesInfo := &types.Info{
		Types:  make(map[ast.Expr]types.TypeAndValue),
		Defs:   make(map[*ast.Ident]types.Object),
		Uses:   make(map[*ast.Ident]types.Object),
		Scopes: make(map[ast.Node]*types.Scope),
	}
	config := &types.Config{
		Importer: importer.Default(),
	}
	if _, err := config.Check(importPath, fset, files, typesInfo); err != nil {
		panic(err)
	}

	for i, file := range files {
		simplifiedFile := astrewrite.Simplify(file, typesInfo, false)
		out, err := os.Create(filepath.Join("goroot", "src", importPath, pkg.GoFiles[i]))
		if err != nil {
			panic(err)
		}
		if err := printer.Fprint(out, fset, simplifiedFile); err != nil {
			panic(err)
		}
		out.Close()
	}
}
