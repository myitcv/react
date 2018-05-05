package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"
)

type coreGen struct {
	buf  *bytes.Buffer
	tbuf *bytes.Buffer
}

func newCoreGen() *coreGen {
	return &coreGen{
		buf:  bytes.NewBuffer(nil),
		tbuf: bytes.NewBuffer(nil),
	}
}

func (c *coreGen) pf(format string, vals ...interface{}) {
	fmt.Fprintf(c.buf, format, vals...)
}

func (c *coreGen) tpf(format string, vals ...interface{}) {
	fmt.Fprintf(c.tbuf, format, vals...)
}

func (c *coreGen) pln(vals ...interface{}) {
	fmt.Fprintln(c.buf, vals...)
}

func (c *coreGen) tpln(vals ...interface{}) {
	fmt.Fprintln(c.tbuf, vals...)
}

func (c *coreGen) pt(tmpl string, val interface{}) {
	tmplExec(c.buf, tmpl, val)
}

func (c *coreGen) tpt(tmpl string, val interface{}) {
	tmplExec(c.tbuf, tmpl, val)
}

func tmplExec(w io.Writer, tmpl string, val interface{}) {
	tmpl = strings.TrimPrefix(tmpl, "\n")

	t := template.New("tmp")

	_, err := t.Parse(tmpl)
	if err != nil {
		fatalf("unable to parse template: %v", err)
	}

	err = t.Execute(w, val)
	if err != nil {
		fatalf("cannot execute template: %v", err)
	}
}
