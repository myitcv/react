package main

import (
	"go/ast"
	"strings"
	"text/template"

	"myitcv.io/immutable"
	"myitcv.io/immutable/util"
)

type immMap struct {
	commonImm

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
			VarName string
			KeyType string
			ValType string
		}{
			Name:    m.name,
			VarName: genVarName(m.name),
			KeyType: o.exprString(m.keyTyp),
			ValType: o.exprString(m.valTyp),
		}

		exp := exporter(m.name)

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
		tmpl.Funcs(exp)
		_, err := tmpl.Parse(immMapTmpl)
		if err != nil {
			fatalf("failed to parse immutable map template: %v", err)
		}

		err = tmpl.Execute(o.output, blanks)
		if err != nil {
			fatalf("failed to execute immutable map template: %v", err)
		}

		o.pt(`
		func (s *{{.}}) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
			if s == nil {
				return true
			}

			if s.Mutable() {
				return false
			}
		`, exp, m.name)

		ktyp := o.exprString(m.keyTyp)
		vtyp := o.exprString(m.valTyp)

		keyIsImm := o.immTypes[strings.TrimPrefix(ktyp, "*")]
		valIsImm := o.immTypes[strings.TrimPrefix(vtyp, "*")]

		if keyIsImm == nil {
			i, err := util.IsImmTypeAst(m.keyTyp, m.file.Imports, m.pkg)
			if err != nil {
				fatalf("failed to check IsImmTypeAst: %v", err)
			}
			keyIsImm = i
		}

		if valIsImm == nil {
			i, err := util.IsImmTypeAst(m.valTyp, m.file.Imports, m.pkg)
			if err != nil {
				fatalf("failed to check IsImmTypeAst: %v", err)
			}
			valIsImm = i
		}

		keyIsImmOk := false
		switch keyIsImm.(type) {
		case util.ImmTypeAstSlice, util.ImmTypeAstStruct, util.ImmTypeAstMap,
			util.ImmTypeAstImplsIntf, util.ImmTypeAstExtIntf:
			keyIsImmOk = true
		}

		valIsImmOk := false
		switch valIsImm.(type) {
		case util.ImmTypeAstSlice, util.ImmTypeAstStruct, util.ImmTypeAstMap,
			util.ImmTypeAstImplsIntf, util.ImmTypeAstExtIntf:
			valIsImmOk = true
		}

		if keyIsImmOk || valIsImmOk {
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

			`, exp, m.name)

			switch {
			case keyIsImmOk && valIsImmOk:
				o.pt(`
				for k, v := range s.theMap {
				`, exp, m.name)
			case keyIsImmOk:
				o.pt(`
				for k := range s.theMap {
				`, exp, m.name)
			case valIsImmOk:
				o.pt(`
				for _, v := range s.theMap {
				`, exp, m.name)
			}

			if keyIsImmOk {
				if _, ok := keyIsImm.(util.ImmTypeAstExtIntf); ok {
					o.pt(`
					switch k.(type) {
					case immutable.Immutable:
						if !k.IsDeeplyNonMutable(seen) {
							return false
						}
					}
					`, exp, m.name)
				} else {
					o.pt(`
					if k != nil && !k.IsDeeplyNonMutable(seen) {
						return false
					}
					`, exp, m.name)
				}
			}

			if valIsImmOk {
				if _, ok := valIsImm.(util.ImmTypeAstExtIntf); ok {
					o.pt(`
					switch v.(type) {
					case immutable.Immutable:
						if !v.IsDeeplyNonMutable(seen) {
							return false
						}
					}
					`, exp, m.name)
				} else {
					o.pt(`
					if v != nil && !v.IsDeeplyNonMutable(seen) {
						return false
					}
					`, exp, m.name)
				}
			}

			o.pt(`
			}
			`, exp, m.name)
		}

		o.pt(`
			return true
		}
		`, exp, m.name)
	}
}
