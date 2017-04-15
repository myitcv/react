package main

import (
	"go/ast"
	"go/token"
	"strings"

	"myitcv.io/immutable"
	"myitcv.io/immutable/util"
)

type commonImm struct {
	fset *token.FileSet

	// the full package import path (not just the name)
	// declaring the type
	pkg string

	// the declaring file
	file *ast.File
}

type immStruct struct {
	commonImm

	name string
	dec  *ast.GenDecl
	st   *ast.StructType

	special bool
}

func (o *output) genImmStructs(structs []immStruct) {
	type genField struct {
		Name  string
		Type  string
		f     *ast.Field
		IsImm util.ImmTypeAst
	}

	for _, s := range structs {

		o.printCommentGroup(s.dec.Doc)
		o.printImmPreamble(s.name, s.st)

		// start of struct
		o.pfln("type %v struct {", s.name)

		o.printLeadSpecCommsFor(s.st)

		o.pln("")

		var fields []genField

		for _, f := range s.st.Fields.List {
			names := ""
			sep := ""

			var isImm util.ImmTypeAst
			typ := o.exprString(f.Type)

			isImm = o.immTypes[strings.TrimPrefix(typ, "*")]

			if isImm == nil {
				i, err := util.IsImmTypeAst(f.Type, s.file.Imports, s.pkg)
				if err != nil {
					panic(err)
				}
				isImm = i
			}

			tag := ""

			if f.Tag != nil {
				tag = f.Tag.Value
			}

			if len(f.Names) == 0 {
				n := strings.TrimPrefix(typ, "*")
				ps := strings.Split(n, ".")
				n = ps[len(ps)-1]
				names = fieldHidingPrefix + n

				fields = append(fields, genField{
					Name:  n,
					Type:  typ,
					f:     f,
					IsImm: isImm,
				})
			} else {
				for _, n := range f.Names {
					names = names + sep + fieldHidingPrefix + n.Name
					sep = ", "

					fields = append(fields, genField{
						Name:  n.Name,
						Type:  typ,
						f:     f,
						IsImm: isImm,
					})
				}
			}
			o.pfln("%v %v %v", names, typ, tag)
		}

		o.pln("")
		o.pln("mutable bool")
		o.pfln("__tmpl %v%v", immutable.ImmTypeTmplPrefix, s.name)

		// end of struct
		o.pfln("}")

		o.pln()

		o.pfln("var _ immutable.Immutable = new(%v)", s.name)
		o.pfln("var _ = new(%v).__tmpl", s.name)
		o.pln()

		exp := exporter(s.name)

		o.pt(`
		func (s *{{.}}) AsMutable() *{{.}} {
			if s.Mutable() {
				return s
			}

			res := *s
		`, exp, s.name)
		if s.special {
			o.pt(`
			res._Key.Version++
			`, exp, nil)
		}
		o.pt(`
			res.mutable = true
			return &res
		}

		func (s *{{.}}) AsImmutable(v *{{.}}) *{{.}} {
			if s == nil {
				return nil
			}

			if s == v {
				return s
			}

			s.mutable = false
			return s
		}

		func (s *{{.}}) Mutable() bool {
			return s.mutable
		}

		func (s *{{.}}) WithMutable(f func(si *{{.}})) *{{.}} {
			res := s.AsMutable()
			f(res)
			res = res.AsImmutable(s)

			return res
		}

		func (s *{{.}}) WithImmutable(f func(si *{{.}})) *{{.}} {
			prev := s.mutable
			s.mutable = false
			f(s)
			s.mutable = prev

			return s
		}
		`, exp, s.name)

		o.pt(`
		func (s *{{.}}) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
			if s == nil {
				return true
			}

			if s.Mutable() {
				return false
			}

			if seen == nil {
				return s.IsDeeplyNonMutable(make(map[interface{}]bool))
			}

			if seen[s] {
				return true
			}

			seen[s] = true
		`, exp, s.name)

		for _, f := range fields {
			switch f.IsImm.(type) {
			case util.ImmTypeAstSlice, util.ImmTypeAstStruct, util.ImmTypeAstMap, util.ImmTypeAstImplsIntf:

				tmpl := struct {
					TypeName string
					Field    genField
				}{
					TypeName: s.name,
					Field:    f,
				}

				o.pt(`
				{
					v := s.`+fieldHidingPrefix+`{{.Field.Name}}

					if v != nil && !v.IsDeeplyNonMutable(seen) {
						return false
					}
				}
				`, exp, tmpl)
			case util.ImmTypeAstExtIntf:

				tmpl := struct {
					TypeName string
					Field    genField
				}{
					TypeName: s.name,
					Field:    f,
				}

				o.pt(`
				{
					v := s.`+fieldHidingPrefix+`{{.Field.Name}}

					switch v := v.(type) {
					case immutable.Immutable:
						if !v.IsDeeplyNonMutable(seen) {
							return false
						}
					}
				}
				`, exp, tmpl)
			}
		}

		o.pt(`
			return true
		}
		`, exp, s.name)

		for _, f := range fields {
			tmpl := struct {
				TypeName string
				Field    genField
			}{
				TypeName: s.name,
				Field:    f,
			}

			exp := exporter(f.Name)

			o.printCommentGroup(f.f.Doc)

			o.pt(`
			func (s *{{.TypeName}}) {{.Field.Name}}() {{.Field.Type}} {
				return s.`+fieldHidingPrefix+`{{.Field.Name}}
			}

			// {{Export "Set"}}{{Capitalise .Field.Name}} is the setter for {{Capitalise .Field.Name}}()
			func (s *{{.TypeName}}) {{Export "Set"}}{{Capitalise .Field.Name}}(n {{.Field.Type}}) *{{.TypeName}} {
				if s.mutable {
					s.`+fieldHidingPrefix+`{{.Field.Name}} = n
					return s
				}

				res := *s
			`, exp, tmpl)
			if s.special {
				o.pt(`
				res._Key.Version++
				`, exp, tmpl)
			}
			o.pt(`
				res.`+fieldHidingPrefix+`{{.Field.Name}} = n
				return &res
			}
			`, exp, tmpl)
		}
	}
}

func (o *output) printLeadSpecCommsFor(st *ast.StructType) {

	var end token.Pos

	// we are looking for comments before the first field (if there is one)

	if f := st.Fields; f != nil && len(f.List) > 0 {
		end = f.List[0].End()
	} else {
		end = st.End()
	}

	for _, cg := range o.curFile.Comments {
		if cg.Pos() > st.Pos() && cg.End() < end {
			for _, c := range cg.List {
				if strings.HasPrefix(c.Text, "//") && !strings.HasPrefix(c.Text, "// ") {
					o.pln(c.Text)
				}
			}
		}
	}

}
