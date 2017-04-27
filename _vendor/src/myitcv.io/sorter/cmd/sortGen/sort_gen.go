// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main // import "myitcv.io/sorter/cmd/sortGen"

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/imports"

	"myitcv.io/gogenerate"
	"myitcv.io/immutable/util"
	"myitcv.io/sorter"
)

const (
	sortGenCmd  = "sortGen"
	orderPrefix = "order"
)

// matching related vars
var (
	orderFnRegex    *regexp.Regexp
	lowerOrder      string
	upperOrder      string
	invalidFileChar *regexp.Regexp
)

// flags
var (
	fLicenseFile = gogenerate.LicenseFileFlag()
	fGoGenLog    = gogenerate.LogFlag()
)

func init() {
	r, n := utf8.DecodeRuneInString(orderPrefix)
	if r == utf8.RuneError {
		fatalf("OrderPrefix not a UTF8 string?")
	}

	l := string(unicode.ToLower(r))
	u := string(unicode.ToUpper(r))

	suffix := orderPrefix[n:]

	lowerOrder = l + suffix
	upperOrder = u + suffix

	orderFunctionPattern := `^[` + l + u + `]` + suffix + `[[:word:]]+`
	orderFnRegex = regexp.MustCompile(orderFunctionPattern)

	invalidFileChar = regexp.MustCompile(`[[:^word:]]`)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(sortGenCmd + ": ")

	flag.Parse()

	gogenerate.DefaultLogLevel(fGoGenLog, gogenerate.LogFatal)

	envFileName, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	wd, err := os.Getwd()
	if err != nil {
		fatalf("unable to get working directory: %v", err)
	}

	// are we running against the first file that contains the sortGen directive?
	// if not return
	dirFiles, err := gogenerate.FilesContainingCmd(wd, sortGenCmd)
	if err != nil {
		fatalf("could not determine if we are the first file: %v", err)
	}

	if len(dirFiles) == 0 {
		fatalf("cannot find any files containing the %v directive", sortGenCmd)
	}

	if envFileName != dirFiles[0] {
		return
	}

	license, err := gogenerate.CommentLicenseHeader(fLicenseFile)
	if err != nil {
		fatalf("could not comment license file: %v", err)
	}

	// if we get here, we know we are the first file...

	gen(wd, license)
}

func gen(dir, license string) {
	bpkg, err := build.ImportDir(dir, 0)
	if err != nil {
		fatalf("could not load package in dir %v: %v", dir, err)
	}

	g := &generator{
		pkg:      bpkg,
		pkgCache: map[string]*build.Package{bpkg.ImportPath: bpkg},
		typCache: make(map[string]map[string]bool),
		license:  license,
	}

	g.fset = token.NewFileSet()

	// do this early
	g.findImmSlices(g.pkg)

	notGenByUs := func(fi os.FileInfo) bool {
		return !gogenerate.FileGeneratedBy(fi.Name(), sortGenCmd)
	}

	pkgs, err := parser.ParseDir(g.fset, g.pkg.Dir, notGenByUs, 0)
	if err != nil {
		fatalf("could not parse dir %v: %v", g.pkg.Dir, err)
	}

	// we are only interested in the package itself, not the exported test package (xtest)
	pkg, ok := pkgs[g.pkg.Name]
	if !ok {
		fatalf("unable to resolve parsed pkg %v", g.pkg.Name)
	}

	for _, f := range pkg.Files {
		g.buf = bytes.NewBuffer(nil)
		g.file = f

		matches := g.getMatches()

		if len(matches) == 0 {
			continue
		}

		toGen, importMap := g.createToGen(matches)

		if len(toGen) > 0 {
			g.genMatches(toGen, importMap)
		}
	}
}

// a generator is the generator for a given package
type generator struct {
	// the package in which we are generating
	pkg *build.Package

	license string

	fset *token.FileSet

	// a cache of build packages that maps a package import path to
	// its corresponding *build.Package
	pkgCache map[string]*build.Package

	// a cache within this generator instances that maps
	// packages import path to the list of immutable slices defined within
	// that package (that is the type name, not the imm type templ
	// name)
	typCache map[string]map[string]bool

	// the current file being analysed
	file *ast.File

	// current gen output buffer
	buf *bytes.Buffer
}

type toGen struct {
	orderFn string
	typ     string

	recvVar string
	recvTyp string

	imm bool
}

func (g *generator) findImmSlices(bp *build.Package) map[string]bool {
	fset := token.NewFileSet()

	// we are interested in Go files and test files; we won't get the resolution exactly right here
	// in the case that someone is referring to a test type outside the package, but the compiler
	// will catch that case
	interestFile := func(fi os.FileInfo) bool {
		var toCheck []string
		toCheck = append(toCheck, bp.GoFiles...)
		toCheck = append(toCheck, bp.TestGoFiles...)

		for _, v := range toCheck {
			if fi.Name() == v {
				return true
			}
		}

		return false
	}

	pkgs, err := parser.ParseDir(fset, bp.Dir, interestFile, 0)
	if err != nil {
		fatalf("could not parse package %v in dir %v: %v", bp.ImportPath, bp.Dir, err)
	}

	pkg, ok := pkgs[bp.Name]
	if !ok {
		fatalf("could not resolved parsed package %v in dir %v", bp.Name, bp.Dir)
	}

	immSlices := make(map[string]bool)

	for _, f := range pkg.Files {

	Decls:
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue Decls
			}

			if gd.Tok != token.TYPE {
				continue Decls
			}

		Specs:
			for _, s := range gd.Specs {
				ts := s.(*ast.TypeSpec)

				name, isit := util.IsImmTmplAst(ts)

				if !isit {
					continue Specs
				}

				at, ok := ts.Type.(*ast.ArrayType)

				if !ok {
					continue Specs
				}

				if at.Len != nil {
					// it's an array and not a slice
					continue Specs
				}

				immSlices[name] = true
			}
		}
	}

	g.typCache[bp.ImportPath] = immSlices

	return immSlices
}

// getMatches returns a map[string]fileMatches where the string is the file path where
// the matches were found
func (g *generator) getMatches() []match {

	// see whether it imports the sorter package
	theImport := ""

	for _, s := range g.file.Imports {
		if s.Path.Value == `"`+sorter.PkgName+`"` {
			if s.Name != nil {
				theImport = s.Name.Name
			} else {
				naked := strings.Trim(s.Path.Value, `"`)
				parts := strings.Split(naked, "/")
				theImport = parts[len(parts)-1]
			}
		}
	}

	if theImport == "" {
		// if it's not import we can't have any templates in here...
		return nil
	}

	var matches []match

Decls:
	for _, d := range g.file.Decls {
		fun, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fn := fun.Name.Name

		if !orderFnRegex.MatchString(fn) {
			continue
		}

		if fun.Type.Results == nil || len(fun.Type.Results.List) != 1 {
			continue
		}

		typ, ok := fun.Type.Results.List[0].Type.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		id, ok := typ.X.(*ast.Ident)
		if !ok {
			continue
		}

		if id.Name != theImport {
			continue
		}

		if typ.Sel.Name != sorter.OrderedName {
			continue
		}

		if fun.Type.Params == nil {
			continue
		}

		// we need to gather the number of params....
		var paramList []ast.Expr
		for _, f := range fun.Type.Params.List {
			for _ = range f.Names {
				paramList = append(paramList, f.Type)
			}
		}

		if len(paramList) != 3 {
			continue
		}

		m := match{
			fun: fun,
		}

		if g.isSliceExpr(paramList[0]) {
			m.orderTyp = paramList[0]
			m.isImmSlice = false
		}

		if g.isImmSliceExpr(paramList[0]) {
			m.orderTyp = paramList[0]
			m.isImmSlice = true
		}

		if m.orderTyp == nil {
			continue
		}

		for i := 1; i < len(paramList); i++ {
			if id, ok := paramList[i].(*ast.Ident); !ok || id.Name != "int" {
				continue Decls
			}
		}

		infof("found a match at %v", g.fset.Position(fun.Pos()))

		matches = append(matches, m)
	}

	return matches
}

func (g *generator) createToGen(matches []match) ([]toGen, map[string]bool) {

	var funs []toGen
	importMap := make(map[string]bool)

	// we need to union the list of functions
	for _, match := range matches {
		var buf bytes.Buffer

		if err := printer.Fprint(&buf, g.fset, match.orderTyp); err != nil {
			fatalf("could not ast print type: %v", err)
		}

		sliceIdent := buf.String()

		recv := ""
		recvVar := ""

		if match.fun.Recv != nil {
			var buf bytes.Buffer

			// we know at this point we have a valid method...
			recvVar = match.fun.Recv.List[0].Names[0].Name

			buf.WriteString("(")
			buf.WriteString(recvVar)

			if err := printer.Fprint(&buf, g.fset, match.fun.Recv.List[0].Type); err != nil {
				fatalf("could not ast print recv: %v", err)
			}

			buf.WriteString(")")

			recv = buf.String()
		}

		// we need to calculate the required imports
		importMatches := findImports(match.orderTyp, g.file.Imports)

		for i := range importMatches {
			importName := i.Path.Value
			if i.Name != nil {
				importName = i.Name.Name + " " + importName
			}

			importMap[importName] = true
		}

		funs = append(funs, toGen{
			orderFn: match.fun.Name.Name,
			typ:     sliceIdent,
			recvTyp: recv,
			recvVar: recvVar,

			imm: match.isImmSlice,
		})
	}

	return funs, importMap
}

// a match is a container for a matching func/meth in a file
type match struct {
	// the actual function/method that has matched
	fun *ast.FuncDecl

	// the "type" of the slice parameter (the first one);
	// for a slice this is the ArrayType; for an imm slice
	// this will be a pointer to an imm slice
	orderTyp ast.Expr

	// whether the type is an immutable slice or not (i.e. a regular
	// slice)
	isImmSlice bool
}

func (g *generator) isSliceExpr(e ast.Expr) bool {
	// this type must be either an array type, specifically
	// with Len == nil (which implies a slice)
	at, ok := e.(*ast.ArrayType)
	if !ok || at.Len != nil {
		return false
	}

	return true
}

func (g *generator) isImmSliceExpr(e ast.Expr) bool {
	// must be a pointer to an immutable slice
	// could be defined in this package or another package

	ste, ok := e.(*ast.StarExpr)
	if !ok {
		return false
	}

	var ip string
	var typ string

	if se, ok := ste.X.(*ast.SelectorExpr); ok {
		// X must be simple ident.ident
		// this is safe to do because the definition of order functions can only be top level
		// hence there can be no variable/anything in a package that compiles that shadows and
		// import name; hence a selctor expression where the X matches an import

		pid, ok := se.X.(*ast.Ident)
		if !ok {
			return false
		}

		// resolve this ident to a package import path
		for _, is := range g.file.Imports {

			p := strings.Trim(is.Path.Value, "\"")

			if (is.Name != nil && is.Name.Name == pid.Name) || path.Base(p) == pid.Name {
				ip = p
				break
			}
		}

		typ = se.Sel.Name
	} else if i, ok := ste.X.(*ast.Ident); ok {
		// this package
		ip = g.pkg.ImportPath
		typ = i.Name
	} else {
		return false
	}

	if ip == "" || typ == "" {
		return false
	}

	if matches, ok := g.typCache[ip]; ok {
		return matches[typ]
	}

	bp, ok := g.pkgCache[ip]
	if !ok {
		p, err := build.Import(ip, g.pkg.Dir, 0)
		if err != nil {
			fatalf("could not load package %q with rel dir %q: %v", ip, g.pkg.Dir, err)
		}

		g.pkgCache[ip] = p
		bp = p
	}

	return g.findImmSlices(bp)[typ]
}

func (g *generator) genMatches(funs []toGen, imps map[string]bool) {

	// we need to generate one file for non-test matches... and one for test matches

	fn := g.fset.Position(g.file.Pos()).Filename

	ofName, ok := gogenerate.NameFileFromFile(fn, sortGenCmd)
	if !ok {
		fatalf("could not name generated file from %v", fn)
	}

	g.pf(`// Code generated by sortGen. DO NOT EDIT.

		`)

	g.pf(g.license)

	g.pf(`package %v

			import "sort"
			import "%v"

		`, g.pkg.Name, sorter.PkgName)

	for i := range imps {
		g.pf("import %v\n", i)
	}

	for _, toGen := range funs {
		tmpl := struct {
			Recv  string
			Sort  string
			Name  string
			Typ   string
			Order string
		}{
			Recv:  toGen.recvTyp,
			Typ:   toGen.typ,
			Order: toGen.orderFn,
		}

		if toGen.recvTyp != "" {
			tmpl.Order = toGen.recvVar + "." + toGen.orderFn
		}

		sortFns := sortFunctions(toGen.orderFn)

		for i, sn := range []string{"Sort", "Stable"} {
			tmpl.Sort = sn
			tmpl.Name = sortFns[i]

			if toGen.imm {
				g.pt(`
					func {{.Recv}} {{.Name}}(vs {{.Typ}}) {{.Typ}}{
						theVs := vs.AsMutable()

						sort.{{.Sort}}(&sorter.Wrapper{
							LenFunc: func() int {
								return theVs.Len()
							},
							LessFunc: func(i, j int) bool {
								return bool({{.Order}}(theVs, i, j))
							},
							SwapFunc: func(i, j int) {
								jPrev := theVs.Get(j)
								iPrev := theVs.Get(i)

								theVs.Set(j, iPrev)
								theVs.Set(i, jPrev)
							},
						})

						return theVs.AsImmutable(vs)
					}
					`, tmpl)
			} else {
				g.pt(`
					func {{.Recv}} {{.Name}}(vs {{.Typ}}) {
						sort.{{.Sort}}(&sorter.Wrapper{
							LenFunc: func() int {
								return len(vs)
							},
							LessFunc: func(i, j int) bool {
								return bool({{.Order}}(vs, i, j))
							},
							SwapFunc: func(i, j int) {
								vs[i], vs[j] = vs[j], vs[i]
							},
						})
					}
					`, tmpl)
			}

		}
	}

	toWrite := g.buf.Bytes()

	res, err := imports.Process(ofName, toWrite, nil)
	if err == nil {
		toWrite = res
	}

	wrote, err := gogenerate.WriteIfDiff(toWrite, ofName)
	if err != nil {
		fatalf("could not write %v: %v", ofName, err)
	}

	if wrote {
		infof("writing %v", ofName)
	} else {
		infof("skipping writing of %v; it's identical", ofName)
	}
}

func (g *generator) pf(format string, args ...interface{}) {
	fmt.Fprintf(g.buf, format, args...)
}

func (g *generator) pt(tmpl string, val interface{}) {
	// on the basis most templates are for convenience define inline
	// as raw string literals which start the ` on one line but then start
	// the template on the next (for readability) we strip the first leading
	// \n if one exists
	tmpl = strings.TrimPrefix(tmpl, "\n")

	t := template.New("tmp")

	_, err := t.Parse(tmpl)
	if err != nil {
		fatalf("unable to parse template: %v", err)
	}

	err = t.Execute(g.buf, val)
	if err != nil {
		fatalf("cannot execute template: %v", err)
	}
}

func sortFunctions(orderFn string) []string {
	// TODO this can be improved

	lower := false
	split := ""

	if strings.HasPrefix(orderFn, upperOrder) {
		split = upperOrder
	} else {
		lower = true
		split = lowerOrder
	}

	parts := strings.SplitAfterN(orderFn, split, 2)

	if lower {
		return []string{"sort" + parts[1], "stableSort" + parts[1]}
	}

	return []string{"Sort" + parts[1], "StableSort" + parts[1]}
}

type importFinder struct {
	imports []*ast.ImportSpec
	matches map[*ast.ImportSpec]bool
}

func (i *importFinder) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.SelectorExpr:
		if x, ok := node.X.(*ast.Ident); ok {
			for _, imp := range i.imports {
				if imp.Name != nil {
					if x.Name == imp.Name.Name {
						i.matches[imp] = true
					}
				} else {
					cleanPath := strings.Trim(imp.Path.Value, "\"")
					parts := strings.Split(cleanPath, "/")
					if x.Name == parts[len(parts)-1] {
						i.matches[imp] = true
					}
				}
			}

		}
	}

	return i
}

func findImports(exp ast.Expr, imports []*ast.ImportSpec) map[*ast.ImportSpec]bool {
	finder := &importFinder{
		imports: imports,
		matches: make(map[*ast.ImportSpec]bool),
	}

	ast.Walk(finder, exp)

	return finder.matches
}

func fatalf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func infof(format string, args ...interface{}) {
	if *fGoGenLog == string(gogenerate.LogInfo) {
		log.Printf(format, args...)
	}
}
