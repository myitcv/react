package main

import (
	"go/ast"
	"go/token"
	"text/template"

	"github.com/myitcv/immutable"
)

type immMap struct {
	fset *token.FileSet

	name   string
	dec    *ast.GenDecl
	typ    ast.Expr
	keyTyp ast.Expr
	valTyp ast.Expr
}

func (o *output) genImmMaps(maps []immMap) {
	for _, m := range maps {
		blanks := struct {
			Name    string
			KeyType string
			ValType string
		}{
			Name:    m.name,
			KeyType: o.exprString(m.keyTyp),
			ValType: o.exprString(m.valTyp),
		}

		fm := exporter(m.name)

		o.printCommentGroup(m.dec.Doc)
		o.printImmPreamble(m.name, m.typ)

		// start of struct
		o.pfln("type %v struct {", m.name)
		o.pln("")

		o.pfln("theMap map[%v]%v", blanks.KeyType, blanks.ValType)
		o.pln("mutable bool")
		o.pfln("__tmpl %v%v", immutable.ImmTypeTmplPrefix, m.name)

		// end of struct
		o.pfln("}")

		tmpl := template.New("immmap")
		tmpl.Funcs(fm)
		_, err := tmpl.Parse(immMapTmpl)
		if err != nil {
			fatalf("failed to parse immutable map template: %v", err)
		}

		err = tmpl.Execute(o.output, blanks)
		if err != nil {
			fatalf("failed to execute immutable map template: %v", err)
		}
	}
}
