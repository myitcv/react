// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

const immMapTmpl = `
var _ immutable.Immutable = new({{.Name}})
var _ = new({{.Name}}).__tmpl

func {{Export "New"}}{{Capitalise .Name}}(inits ...func(m *{{.Name}})) *{{.Name}} {
	res := {{Export "New"}}{{Capitalise .Name}}Cap(0)
	if len(inits) == 0 {
		return res
	}

	return res.WithMutable(func (m *{{.Name}}) {
		for _, i := range inits {
			i(m)
		}
	})
}

func {{Export "New"}}{{Capitalise .Name}}Cap(l int) *{{.Name}} {
	return &{{.Name}}{
		theMap: make(map[{{.KeyType}}]{{.ValType}}, l),
	}
}

func (m *{{.Name}})Mutable() bool {
	return m.mutable
}

func (m *{{.Name}}) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theMap)
}

func (m *{{.Name}}) Get(k {{.KeyType}}) ({{.ValType}}, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m *{{.Name}}) AsMutable() *{{.Name}} {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *{{.Name}}) dup() *{{.Name}} {
	resMap := make(map[{{.KeyType}}]{{.ValType}}, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := &{{.Name}}{
		theMap: resMap,
	}

	return res
}

func (m *{{.Name}}) AsImmutable(v *{{.Name}}) *{{.Name}} {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *{{.Name}}) Range() map[{{.KeyType}}]{{.ValType}} {
	if m == nil {
		return nil
	}

	return m.theMap
}

func (mr *{{.Name}}) WithMutable(f func({{.VarName}} *{{.Name}})) *{{.Name}} {
	res := mr.AsMutable()
	f(res)
	res = res.AsImmutable(mr)

	return res
}

func (mr *{{.Name}}) WithImmutable(f func({{.VarName}} *{{.Name}})) *{{.Name}} {
	prev := mr.mutable
	mr.mutable = false
	f(mr)
	mr.mutable = prev

	return mr
}

func (m *{{.Name}}) Set(k {{.KeyType}}, v {{.ValType}}) *{{.Name}} {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	res := m.dup()
	res.theMap[k] = v

	return res
}

func (m *{{.Name}}) Del(k {{.KeyType}}) *{{.Name}} {
	if _, ok := m.theMap[k]; !ok {
		return m
	}

	if m.mutable {
		delete(m.theMap, k)
		return m
	}

	res := m.dup()
	delete(res.theMap, k)

	return res
}
`
