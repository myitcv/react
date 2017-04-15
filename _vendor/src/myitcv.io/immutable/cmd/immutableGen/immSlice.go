package main

import (
	"go/ast"
	"strings"
	"text/template"

	"myitcv.io/immutable"
	"myitcv.io/immutable/util"
)

type immSlice struct {
	commonImm

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

		exp := exporter(s.name)

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
		tmpl.Funcs(exp)
		_, err := tmpl.Parse(immSliceTmpl)
		if err != nil {
			fatalf("failed to parse immutable slice template: %v", err)
		}

		err = tmpl.Execute(o.output, blanks)
		if err != nil {
			fatalf("failed to execute immutable slice template: %v", err)
		}

		o.pt(`
		func (s *{{.}}) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
			if s == nil {
				return true
			}

			if s.Mutable() {
				return false
			}
		`, exp, s.name)

		vtyp := o.exprString(s.valTyp)

		valIsImm := o.immTypes[strings.TrimPrefix(vtyp, "*")]

		if valIsImm == nil {
			i, err := util.IsImmTypeAst(s.valTyp, s.file.Imports, s.pkg)
			if err != nil {
				fatalf("failed to check IsImmTypeAst: %v", err)
			}
			valIsImm = i
		}

		switch valIsImm.(type) {
		case util.ImmTypeAstSlice, util.ImmTypeAstStruct, util.ImmTypeAstMap, util.ImmTypeAstImplsIntf:
			o.pt(`
			if s.Len() == 0 {
				return true
			}

			if seen == nil {
				return s.IsDeeplyNonMutable(make(map[interface{}]bool))
			}

			if seen[s] {
				return true
			}

			seen[s] = true

			for _, v := range s.theSlice {
			`, exp, s.name)

			if _, ok := valIsImm.(util.ImmTypeAstExtIntf); ok {
				o.pt(`
					switch v := v.(type) {
					case immutable.Immutable:
						if !v.IsDeeplyNonMutable(seen) {
							return false
						}
					}
				`, exp, s.name)
			} else {
				o.pt(`
					if v != nil && !v.IsDeeplyNonMutable(seen) {
						return false
					}
				`, exp, s.name)
			}

			o.pt(`
			}
			`, exp, s.name)
		}

		o.pt(`
			return true
		}
		`, exp, s.name)
	}
}
