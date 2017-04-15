package main

import (
	"go/ast"
	"strings"
)

type importFinder struct {
	imports []*ast.ImportSpec
	matches map[*ast.ImportSpec]struct{}
}

func (i *importFinder) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.SelectorExpr:
		if x, ok := node.X.(*ast.Ident); ok {
			for _, imp := range i.imports {
				if imp.Name != nil {
					if x.Name == imp.Name.Name {
						i.matches[imp] = struct{}{}
					}
				} else {
					cleanPath := strings.Trim(imp.Path.Value, "\"")
					parts := strings.Split(cleanPath, "/")
					if x.Name == parts[len(parts)-1] {
						i.matches[imp] = struct{}{}
					}
				}
			}
		}
	}
	return i
}
