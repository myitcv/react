package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

type field struct {
	Name string
	Type string

	Omit bool
}

type fieldExploder struct {
	first bool

	pkgStr string
	sn     string

	fields []field
	imps   map[*ast.ImportSpec]struct{}
}

func (fe *fieldExploder) explode() error {
	tf, err := loadStructType(fe.pkgStr, fe.sn)
	if err != nil {
		return err
	}

	if tf == nil {
		// we can't do any better...
		return nil
	}

	st := tf.ts.Type.(*ast.StructType)

	for ind, f := range st.Fields.List {
		if f.Names == nil {
			// embedded struct
			// cannot be a paren expr

			t := f.Type

			if v, ok := t.(*ast.StarExpr); ok {
				t = v.X
			}

			var nfe *fieldExploder

			switch v := t.(type) {
			case *ast.Ident:
				// same package

				nfe = &fieldExploder{
					first:  fe.first && ind == 0,
					pkgStr: fe.pkgStr,
					sn:     v.Name,
					imps:   fe.imps,
				}

			case *ast.SelectorExpr:
				x, ok := v.X.(*ast.Ident)
				if !ok {
					s := astNodeString(f.Type)
					fatalf("don't know how to handle a field type: %v", s)
				}

				typStr := v.Sel.Name

				var pkgStr string

				foundImp := false

				for _, i := range tf.file.Imports {
					p := strings.Trim(i.Path.Value, "\"")

					toCheck := path.Base(p)

					if i.Name != nil {
						toCheck = i.Name.Name
					}

					if x.Name == toCheck {
						pkgStr = p
						foundImp = true

						break
					}
				}

				if !foundImp {
					// let's move on for now
					continue
				}

				nfe = &fieldExploder{
					first:  fe.first && ind == 0,
					pkgStr: pkgStr,
					sn:     typStr,
					imps:   fe.imps,
				}
			}

			err := nfe.explode()
			if err != nil {
				return err
			}

			fe.fields = append(fe.fields, nfe.fields...)

		} else {
			// real fields - we can print the type... but need
			// to collect the imports first

			newImps := make(map[*ast.ImportSpec]struct{})

			i := &importFinder{
				imports: tf.file.Imports,
				matches: newImps,
			}

			ast.Walk(i, f.Type)

			ts := astNodeString(f.Type)

			var omit bool

			if f.Tag != nil {
				omit = strings.Contains(f.Tag.Value, `react:"omitempty"`)
			}

			if ind == 0 && fe.first && i.isJs {
				continue
			}

			for k := range newImps {
				fe.imps[k] = struct{}{}
			}

			for _, n := range f.Names {
				fe.fields = append(fe.fields, field{
					Name: n.Name,
					Type: ts,
					Omit: omit,
				})
			}
		}
	}

	return nil
}

var pkgCache = make(map[string]*ast.Package)

func loadStructType(pkgStr string, sn string) (*typeFile, error) {
	pkg, err := loadPkg(pkgStr)
	if err != nil {
		return nil, err
	}

	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			if gd.Tok != token.TYPE {
				continue
			}

			for _, s := range gd.Specs {
				ts := s.(*ast.TypeSpec)

				if _, ok := ts.Type.(*ast.StructType); ok && ts.Name.Name == sn {
					return &typeFile{
						file: f,
						ts:   ts,
					}, nil
				}
			}

		}
	}

	return nil, nil
}

func loadPkg(pkgStr string) (*ast.Package, error) {
	pkg, ok := pkgCache[pkgStr]
	if !ok {
		p, err := loadPkgImpl(pkgStr)
		if err != nil {
			return nil, fmt.Errorf("failed to load package %v: %v", pkgStr, err)
		}

		pkg = p
		pkgCache[pkgStr] = p
	}

	return pkg, nil
}

func loadPkgImpl(pkgStr string) (*ast.Package, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not get working directory: %v", err)
	}

	bpkg, err := build.Import(pkgStr, wd, 0)
	if err != nil {
		return nil, fmt.Errorf("could not resolve %v: %v", pkgStr, err)
	}

	pkgs, err := parser.ParseDir(fset, bpkg.Dir, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("could not parse %v in %v: %v", pkgStr, bpkg.Dir, err)
	}

	base := path.Base(pkgStr)

	p, ok := pkgs[base]
	if !ok {
		return nil, fmt.Errorf("failed to find package %v (%v) in %v", base, pkgStr, bpkg.Dir)
	}

	return p, nil
}

type importFinder struct {
	imports []*ast.ImportSpec
	matches map[*ast.ImportSpec]struct{}

	isJs bool
}

func (i *importFinder) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.SelectorExpr:
		if x, ok := node.X.(*ast.Ident); ok {
			for _, imp := range i.imports {

				var toCheck string
				cleanPath := strings.Trim(imp.Path.Value, "\"")

				if imp.Name != nil {
					toCheck = imp.Name.Name
				} else {
					toCheck = path.Base(cleanPath)
				}

				if x.Name == toCheck {
					i.isJs = cleanPath == jsPkg
					i.matches[imp] = struct{}{}
				}
			}
		}
	}

	return i
}
