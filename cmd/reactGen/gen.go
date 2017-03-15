// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/imports"

	"github.com/myitcv/gogenerate"
)

const (
	reactPkg      = "github.com/myitcv/gopherjs/react"
	compDefName   = "ComponentDef"
	compDefSuffix = "Def"

	stateTypeSuffix = "State"
	propsTypeSuffix = "Props"

	getInitialState           = "GetInitialState"
	componentWillReceiveProps = "ComponentWillReceiveProps"
	equals                    = "Equals"
)

type gen struct {
	fset *token.FileSet

	pkg string

	components    map[string]*ast.TypeSpec
	types         map[string]*ast.TypeSpec
	pointMeths    map[string][]*ast.FuncDecl
	nonPointMeths map[string][]*ast.FuncDecl
}

func dogen(dir, license string) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		fatalf("unable to parse %v: %v", dir, err)
	}

	// we intentionally walk all packages, i.e. the package in the current directory
	// and any x-test package that may also be present
	for pn, pkg := range pkgs {
		g := &gen{
			fset: fset,
			pkg:  pn,

			components:    make(map[string]*ast.TypeSpec),
			types:         make(map[string]*ast.TypeSpec),
			pointMeths:    make(map[string][]*ast.FuncDecl),
			nonPointMeths: make(map[string][]*ast.FuncDecl),
		}

		for fn, file := range pkg.Files {

			if gogenerate.FileGeneratedBy(fn, "reactGen") {
				continue
			}

			foundImp := false
			impName := ""

			for _, i := range file.Imports {
				p := strings.Trim(i.Path.Value, "\"")

				if p == reactPkg {
					foundImp = true

					if i.Name != nil {
						impName = i.Name.Name
					} else {
						impName = path.Base(reactPkg)
					}

					break
				}
			}

			if !foundImp {
				continue
			}

			for _, d := range file.Decls {
				switch d := d.(type) {
				case *ast.FuncDecl:
					if d.Recv == nil {
						continue
					}

					f := d.Recv.List[0]

					switch v := f.Type.(type) {
					case *ast.StarExpr:
						id, ok := v.X.(*ast.Ident)
						if !ok {
							continue
						}
						g.pointMeths[id.Name] = append(g.pointMeths[id.Name], d)
					case *ast.Ident:
						g.nonPointMeths[v.Name] = append(g.pointMeths[v.Name], d)
					}

				case *ast.GenDecl:
					if d.Tok != token.TYPE {
						continue
					}

					for _, ts := range d.Specs {
						ts := ts.(*ast.TypeSpec)

						st, ok := ts.Type.(*ast.StructType)
						if !ok || st.Fields == nil {
							continue
						}

						foundAnon := false

						for _, f := range st.Fields.List {
							if f.Names != nil {
								// it must be anonymous
								continue
							}

							se, ok := f.Type.(*ast.SelectorExpr)
							if !ok {
								continue
							}

							if se.Sel.Name != compDefName {
								continue
							}

							id, ok := se.X.(*ast.Ident)
							if !ok {
								continue
							}

							if id.Name != impName {
								continue
							}

							foundAnon = true
						}

						if foundAnon && strings.HasSuffix(ts.Name.Name, compDefSuffix) {
							g.components[ts.Name.Name] = ts
						} else {
							g.types[ts.Name.Name] = ts
						}
					}
				}
			}
		}

		// at this point we have the components and their methods
		for cd := range g.components {
			g.genComp(cd)
		}
	}
}

type compGen struct {
	*gen

	Recv string
	Name string

	HasState                     bool
	HasProps                     bool
	HasGetInitState              bool
	HasComponentWillReceiveProps bool

	PropsHasEquals bool
	StateHasEquals bool

	buf *bytes.Buffer
}

func (g *gen) genComp(defName string) {

	name := strings.TrimSuffix(defName, compDefSuffix)

	r, _ := utf8.DecodeRuneInString(name)

	cg := &compGen{
		gen:  g,
		buf:  bytes.NewBuffer(nil),
		Name: name,
		Recv: string(unicode.ToLower(r)),
	}

	_, hasState := g.types[name+stateTypeSuffix]
	_, hasProps := g.types[name+propsTypeSuffix]

	cg.HasState = hasState
	cg.HasProps = hasProps

	if hasState {
		for _, m := range g.pointMeths[defName] {
			if m.Name.Name != getInitialState {
				continue
			}

			if m.Type.Params != nil && len(m.Type.Params.List) > 0 {
				continue
			}

			if m.Type.Results != nil && len(m.Type.Results.List) != 1 {
				continue
			}

			rp := m.Type.Results.List[0]

			id, ok := rp.Type.(*ast.Ident)
			if !ok {
				continue
			}

			if id.Name == name+stateTypeSuffix {
				cg.HasGetInitState = true
				break
			}
		}

		for _, m := range g.nonPointMeths[name+stateTypeSuffix] {
			if m.Name.Name != equals {
				continue
			}

			if m.Type.Params != nil && len(m.Type.Params.List) != 1 {
				continue
			}

			if m.Type.Results != nil && len(m.Type.Results.List) != 1 {
				continue
			}

			{
				v := m.Type.Params.List[0]

				id, ok := v.Type.(*ast.Ident)
				if !ok {
					continue
				}

				if id.Name != name+stateTypeSuffix {
					continue
				}
			}

			{
				v := m.Type.Results.List[0]

				id, ok := v.Type.(*ast.Ident)
				if !ok {
					continue
				}

				if id.Name != "bool" {
					continue
				}
			}

			cg.StateHasEquals = true
		}
	}

	if hasProps {
		for _, m := range g.pointMeths[defName] {
			if m.Name.Name != componentWillReceiveProps {
				continue
			}

			if m.Type.Params != nil && len(m.Type.Params.List) != 1 {
				continue
			}

			if m.Type.Results != nil && len(m.Type.Results.List) != 0 {
				continue
			}

			p := m.Type.Params.List[0]

			id, ok := p.Type.(*ast.Ident)
			if !ok {
				continue
			}

			if id.Name == name+propsTypeSuffix {
				cg.HasComponentWillReceiveProps = true
				break
			}
		}

		for _, m := range g.nonPointMeths[name+propsTypeSuffix] {
			if m.Name.Name != equals {
				continue
			}

			if m.Type.Params != nil && len(m.Type.Params.List) != 1 {
				continue
			}

			if m.Type.Results != nil && len(m.Type.Results.List) != 1 {
				continue
			}

			{
				v := m.Type.Params.List[0]

				id, ok := v.Type.(*ast.Ident)
				if !ok {
					continue
				}

				if id.Name != name+propsTypeSuffix {
					continue
				}
			}

			{
				v := m.Type.Results.List[0]

				id, ok := v.Type.(*ast.Ident)
				if !ok {
					continue
				}

				if id.Name != "bool" {
					continue
				}
			}

			cg.PropsHasEquals = true
		}
	}

	cg.pf("// Code generated by %v; DO NOT EDIT.\n", reactGenCmd)
	cg.pln()
	cg.pf("package %v\n", cg.pkg)
	cg.pf("import \"%v\"\n", reactPkg)
	cg.pln()

	cg.pt(`
func ({{.Recv}} *{{.Name}}Def) ShouldComponentUpdateIntf(nextProps interface{}) bool {
	{{if and .HasProps -}}
	{{if .PropsHasEquals -}}
	return {{.Recv}}.Props().Equals(nextProps.({{.Name}}Props))
	{{else -}}
	return {{.Recv}}.Props() == nextProps.({{.Name}}Props)
	{{end -}}
	{{else if .HasState -}}
	return true
	{{else -}}
	// no props or state... so nothing would cause this to require re-rendering
	return false
	{{end -}}
}

{{if .HasState}}
// SetState is an auto-generated proxy proxy to update the state for the
// {{.Name}} component.  SetState does not immediately mutate {{.Recv}}.State()
// but creates a pending state transition.
func ({{.Recv}} *{{.Name}}Def) SetState(s {{.Name}}State) {
	{{.Recv}}.ComponentDef.SetState(s)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the {{.Name}} component
func ({{.Recv}} *{{.Name}}Def) State() {{.Name}}State {
	return {{.Recv}}.ComponentDef.State().({{.Name}}State)
}

// IsState is an auto-generated definition so that {{.Name}}State implements
// the github.com/myitcv/gopherjs/react.State interface.
func ({{.Recv}} {{.Name}}State) IsState() {}

var _ react.State = {{.Name}}State{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func ({{.Recv}} *{{.Name}}Def) GetInitialStateIntf() react.State {
{{if .HasGetInitState -}}
	return {{.Recv}}.GetInitialState()
{{else -}}
	return {{.Name}}State{}
{{end -}}
}

func ({{.Recv}} {{.Name}}State) EqualsIntf(v interface{}) bool {
	{{if .StateHasEquals -}}
	return {{.Recv}}.Equals(v.({{.Name}}State))
	{{else -}}
	return {{.Recv}} == v.({{.Name}}State)
	{{end -}}
}
{{end}}


{{if .HasProps}}
// Props is an auto-generated proxy to the current props of {{.Name}}
func ({{.Recv}} *{{.Name}}Def) Props() {{.Name}}Props {
	uprops := {{.Recv}}.ComponentDef.Props()
	return uprops.({{.Name}}Props)
}

{{if .HasComponentWillReceiveProps}}
// ComponentWillReceivePropsIntf is an auto-generated proxy to
// ComponentWillReceiveProps
func ({{.Recv}} *{{.Name}}Def) ComponentWillReceivePropsIntf(i interface{}) {
	ourProps := i.({{.Name}}Props)
	{{.Recv}}.ComponentWillReceiveProps(ourProps)
}
{{end}}

func ({{.Recv}} {{.Name}}Props) EqualsIntf(v interface{}) bool {
	{{if .PropsHasEquals -}}
	return {{.Recv}}.Equals(v.({{.Name}}Props))
	{{else -}}
	return {{.Recv}} == v.({{.Name}}Props)
	{{end -}}
}

var _ react.Equals = {{.Name}}Props{}
{{end}}
	`, cg)

	ofName := gogenerate.NameFile(name, reactGenCmd)
	toWrite := cg.buf.Bytes()

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

func (c *compGen) pf(format string, vals ...interface{}) {
	fmt.Fprintf(c.buf, format, vals...)
}

func (c *compGen) pln(vals ...interface{}) {
	fmt.Fprintln(c.buf, vals...)
}

func (c *compGen) pt(tmpl string, val interface{}) {
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

	err = t.Execute(c.buf, val)
	if err != nil {
		fatalf("cannot execute template: %v", err)
	}
}
