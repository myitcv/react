// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"myitcv.io/gogenerate"
	"myitcv.io/immutable/util"
)

const (
	fieldHidingPrefix = "_"
)

func execute(dir string, envPkg string, licenseHeader string, cmds gogenCmds) {

	absDir, err := filepath.Abs(dir)
	if err != nil {
		fatalf("could not make absolute path from %v: %v", dir, err)
	}

	bpkg, err := build.ImportDir(absDir, 0)
	if err != nil {
		fatalf("could not resolve package from dir %v: %v", dir, err)
	}

	fset := token.NewFileSet()

	notGenByUs := func(fi os.FileInfo) bool {
		return !gogenerate.FileGeneratedBy(fi.Name(), immutableGenCmd)
	}

	pkgs, err := parser.ParseDir(fset, dir, notGenByUs, parser.AllErrors|parser.ParseComments)
	if err != nil {
		fatalf("could not parse dir %v: %v", dir, err)
	}

	pkg, ok := pkgs[envPkg]

	if !ok {
		pps := make([]string, 0, len(pkgs))
		for k := range pkgs {
			pps = append(pps, k)
		}
		fatalf("expected to have parsed %v, instead parsed %v", envPkg, pps)
	}

	out := &output{
		dir:       dir,
		fset:      fset,
		pkg:       envPkg,
		license:   licenseHeader,
		goGenCmds: cmds,
		files:     make(map[*ast.File]*fileTmpls),
		cms:       make(map[*ast.File]ast.CommentMap),
	}

	allTypes := make(map[string]util.ImmTypeAst)

	for fn, f := range pkg.Files {
		// skip files that we generated
		if gogenerate.FileGeneratedBy(fn, immutableGenCmd) {
			continue
		}

		cm := ast.NewCommentMap(fset, f, f.Comments)
		og := gatherImmTypes(bpkg.ImportPath, fset, f)
		out.files[f] = og

		for _, m := range og.maps {
			allTypes[m.name] = util.ImmTypeAstMap{
				Key:  m.keyTyp,
				Elem: m.valTyp,
			}
		}

		for _, s := range og.slices {
			allTypes[s.name] = util.ImmTypeAstSlice{
				Elem: s.valTyp,
			}
		}

		for _, s := range og.structs {
			allTypes[s.name] = util.ImmTypeAstStruct{}
		}

		out.cms[f] = cm
	}

	out.immTypes = allTypes

	out.genImmTypes()
}

type output struct {
	dir       string
	pkg       string
	fset      *token.FileSet
	license   string
	goGenCmds gogenCmds

	output *bytes.Buffer

	curFile *ast.File

	// a convenience map of all the imm types we will
	// be generating in this package
	immTypes map[string]util.ImmTypeAst

	files map[*ast.File]*fileTmpls
	cms   map[*ast.File]ast.CommentMap
}

type fileTmpls struct {
	imports map[*ast.ImportSpec]struct{}

	maps    []immMap
	slices  []immSlice
	structs []immStruct
}

func gatherImmTypes(pkg string, fset *token.FileSet, file *ast.File) *fileTmpls {
	g := &fileTmpls{
		imports: make(map[*ast.ImportSpec]struct{}),
	}

	impf := &importFinder{
		imports: file.Imports,
		matches: g.imports,
	}

	comm := commonImm{
		fset: fset,
		file: file,
		pkg:  pkg,
	}

	for _, d := range file.Decls {

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

			infof("found immutable declaration at %v", fset.Position(gd.Pos()))

			switch typ := ts.Type.(type) {
			case *ast.MapType:
				g.maps = append(g.maps, immMap{
					commonImm: comm,
					name:      name,
					dec:       gd,
					typ:       typ,
					keyTyp:    typ.Key,
					valTyp:    typ.Value,
				})

				ast.Walk(impf, ts.Type)

			case *ast.ArrayType:
				// TODO support for arrays

				if typ.Len == nil {
					g.slices = append(g.slices, immSlice{
						commonImm: comm,
						name:      name,
						dec:       gd,
						typ:       typ,
						valTyp:    typ.Elt,
					})
				}

				ast.Walk(impf, ts.Type)

			case *ast.StructType:
				g.structs = append(g.structs, immStruct{
					commonImm: comm,
					name:      name,
					dec:       gd,
					st:        typ,
					special:   isSpecialStruct(name, typ),
				})

				ast.Walk(impf, ts.Type)
			}
		}
	}

	return g
}

func isSpecialStruct(name string, st *ast.StructType) bool {
	// work out whether this is a special struct with a Key field
	// pattern is:
	//
	// 1. struct field has a field called Key of type {{.StructName}}Key (non pointer)
	//
	// later checks will include:
	//
	// 2. said type has two fields, Uuid and Version, of type {{.StructName}}Uuid and uint64 respectively
	// 3. the underlying type of {{.StructName}}Uuid is uint64 (we might be able to relax these two
	// two underlying type restrictions)

	if st.Fields == nil {
		return false
	}

	for _, f := range st.Fields.List {
		idt, ok := f.Type.(*ast.Ident)
		if !ok {
			continue
		}

		if idt.Name != name+"Key" {
			continue
		}

		for _, fn := range f.Names {
			if fn.Name == "Key" {
				return true
			}
		}
	}

	return false
}

func (o *output) genImmTypes() {
	for f, v := range o.files {
		o.curFile = f

		if len(v.maps) == 0 && len(v.slices) == 0 && len(v.structs) == 0 {
			continue
		}

		o.output = bytes.NewBuffer(nil)

		o.pfln("// Code generated by %v. DO NOT EDIT.", immutableGenCmd)
		o.pln("")

		o.pf(o.license)

		o.pf("package %v\n", o.pkg)

		// is there a "standard" place for //go:generate comments?
		for _, v := range o.goGenCmds {
			o.pf("//go:generate %v\n", v)
		}

		o.pln("//immutableVet:skipFile")
		o.pln("")

		o.pln("import (")

		o.pln("\"myitcv.io/immutable\"")
		o.pln()

		for i := range v.imports {
			if i.Name != nil {
				o.pfln("%v %v", i.Name.Name, i.Path.Value)
			} else {
				o.pfln("%v", i.Path.Value)
			}
		}

		o.pln(")")

		o.pln("")

		o.genImmMaps(v.maps)
		o.genImmSlices(v.slices)
		o.genImmStructs(v.structs)

		source := o.output.Bytes()

		toWrite := source

		fn := o.fset.Position(f.Pos()).Filename

		// this is the file path
		offn, ok := gogenerate.NameFileFromFile(fn, immutableGenCmd)
		if !ok {
			fatalf("could not name file from %v", fn)
		}

		out := bytes.NewBuffer(nil)
		cmd := exec.Command("gofmt", "-s")
		cmd.Stdin = o.output
		cmd.Stdout = out

		err := cmd.Run()
		if err == nil {
			toWrite = out.Bytes()
		} else {
			infof("failed to format %v: %v", fn, err)
		}

		wrote, err := gogenerate.WriteIfDiff(toWrite, offn)
		if err != nil {
			fatalf("could not write %v: %v", offn, err)
		}

		if wrote {
			infof("writing %v", offn)
		} else {
			infof("skipping writing of %v; it's identical", offn)
		}
	}
}

func (o *output) exprString(e ast.Expr) string {
	var buf bytes.Buffer

	err := printer.Fprint(&buf, o.fset, e)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func (o *output) printCommentGroup(d *ast.CommentGroup) {
	if d != nil {
		for _, c := range d.List {
			o.pfln("%v", c.Text)
		}
	}
}

func (o *output) printImmPreamble(name string, node ast.Node) {
	fset := o.fset

	if st, ok := node.(*ast.StructType); ok {

		// we need to do some manipulation

		buf := bytes.NewBuffer(nil)

		fmt.Fprintf(buf, "struct {\n")

		if st.Fields != nil {
			line := o.fset.Position(st.Fields.List[0].Pos()).Line

			for _, f := range st.Fields.List {
				curLine := o.fset.Position(f.Pos()).Line

				if line != curLine {
					// catch up
					fmt.Fprintln(buf, "")
					line = curLine
				}

				ids := make([]string, 0, len(f.Names))
				for _, n := range f.Names {
					ids = append(ids, n.Name)
				}
				fmt.Fprintf(buf, "%v %v\n", strings.Join(ids, ","), o.exprString(f.Type))

				line++
			}
		}

		fmt.Fprintf(buf, "}")

		exprStr := buf.String()

		fset = token.NewFileSet()
		newnode, err := parser.ParseExprFrom(fset, "", exprStr, 0)
		if err != nil {
			fatalf("could not parse documentation struct from %v: %v", exprStr, err)
		}

		node = newnode
	}

	o.pln("//")
	o.pfln("// %v is an immutable type and has the following template:", name)
	o.pln("//")

	tmplBuf := bytes.NewBuffer(nil)

	err := printer.Fprint(tmplBuf, fset, node)
	if err != nil {
		fatalf("could not printer template declaration: %v", err)
	}

	sc := bufio.NewScanner(tmplBuf)
	for sc.Scan() {
		o.pfln("// \t%v", sc.Text())
	}
	if err := sc.Err(); err != nil {
		fatalf("could not scan printed template: %v", err)
	}

	o.pln("//")
}

func (o *output) pln(i ...interface{}) {
	fmt.Fprintln(o.output, i...)
}

func (o *output) pf(format string, i ...interface{}) {
	fmt.Fprintf(o.output, format, i...)
}

func (o *output) pfln(format string, i ...interface{}) {
	o.pf(format+"\n", i...)
}

func (o *output) pt(tmpl string, fm template.FuncMap, val interface{}) {

	// on the basis most templates are for convenience define inline
	// as raw string literals which start the ` on one line but then start
	// the template on the next (for readability) we strip the first leading
	// \n if one exists
	tmpl = strings.TrimPrefix(tmpl, "\n")

	t := template.New("tmp")
	t.Funcs(fm)

	_, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(o.output, val)
	if err != nil {
		panic(err)
	}
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func infoln(args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Println(args...)
	}
}

func infof(format string, args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Printf(format, args...)
	}
}
