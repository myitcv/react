package main

import (
	"go/ast"
	"go/token"
	"text/template"

	"github.com/myitcv/immutable"
)

type immSlice struct {
	fset *token.FileSet

	name   string
	typ    ast.Expr
	valTyp ast.Expr
	dec    *ast.GenDecl
}

func (o *output) genImmSlices(slices []immSlice) {

	for _, s := range slices {
		blanks := struct {
			Name string
			Type string
		}{
			Name: s.name,
			Type: o.exprString(s.valTyp),
		}

		fm := exporter(s.name)

		o.printCommentGroup(s.dec.Doc)
		o.printImmPreamble(s.name, s.typ)

		// start of struct
		o.pfln("type %v struct {", s.name)
		o.pln("")

		o.pfln("theSlice []%v", blanks.Type)
		o.pln("mutable bool")
		o.pfln("__tmpl %v%v", immutable.ImmTypeTmplPrefix, s.name)

		// end of struct
		o.pfln("}")

		tmpl := template.New("immslice")
		tmpl.Funcs(fm)
		_, err := tmpl.Parse(immSliceTmpl)
		if err != nil {
			fatalf("failed to parse immutable slice template: %v", err)
		}

		err = tmpl.Execute(o.output, blanks)
		if err != nil {
			fatalf("failed to execute immutable slice template: %v", err)
		}
	}
}
