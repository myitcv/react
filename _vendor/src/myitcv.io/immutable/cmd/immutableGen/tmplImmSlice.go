// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

const immSliceTmpl = `
var _ immutable.Immutable = new({{.Name}})
var _ = new({{.Name}}).__tmpl

func {{Export "New"}}{{Capitalise .Name}}(s ...{{.Type}}) *{{.Name}} {
	c := make([]{{.Type}}, len(s))
	copy(c, s)

	return &{{.Name}}{
		theSlice: c,
	}
}

func {{Export "New"}}{{Capitalise .Name}}Len(l int) *{{.Name}} {
	c := make([]{{.Type}}, l)

	return &{{.Name}}{
		theSlice: c,
	}
}

func (m *{{.Name}})Mutable() bool {
	return m.mutable
}

func (m *{{.Name}}) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *{{.Name}}) Get(i int) {{.Type}} {
	return m.theSlice[i]
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
	resSlice := make([]{{.Type}}, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &{{.Name}}{
		theSlice: resSlice,
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

func (m *{{.Name}}) Range() []{{.Type}} {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *{{.Name}}) WithMutable(f func(mi *{{.Name}})) *{{.Name}} {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *{{.Name}}) WithImmutable(f func(mi *{{.Name}})) *{{.Name}} {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *{{.Name}}) Set(i int, v {{.Type}}) *{{.Name}} {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *{{.Name}}) Append(v ...{{.Type}}) *{{.Name}} {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
`
