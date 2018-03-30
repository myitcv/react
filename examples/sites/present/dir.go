// Copyright (c) 2018 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

// This package contains a fair amount of copy-paste from
// https://godoc.org/golang.org/x/tools/cmd/present.  The copyright notice from
// that repo can be seen below, and is linked
// https://github.com/golang/tools/blob/master/LICENSE

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
	"path"
	"path/filepath"

	"honnef.co/go/js/xhr"

	"golang.org/x/tools/present"
)

func isDoc(path string) bool {
	_, ok := contentTemplate[filepath.Ext(path)]
	return ok
}

var (
	// dirListTemplate holds the front page template.
	dirListTemplate *template.Template

	// contentTemplate maps the presentable file extensions to the
	// template to be executed.
	contentTemplate map[string]*template.Template
)

func initTemplates(base string) error {
	var err error

	// Locate the template file.
	actionTmplContents, err := Asset(filepath.Join(base, "templates/action.tmpl"))
	if err != nil {
		return err
	}

	contentTemplate = make(map[string]*template.Template)

	for ext, contentTmpl := range map[string]string{
		".slide":   "slides.tmpl",
		".article": "article.tmpl",
	} {
		contentTmpl = filepath.Join(base, "templates", contentTmpl)

		actionTmpl := present.Template()
		actionTmpl = actionTmpl.Funcs(template.FuncMap{"playable": playable})
		actionTmpl, err = actionTmpl.Parse(string(actionTmplContents))
		if err != nil {
			return err
		}

		contents, err := Asset(contentTmpl)
		if err != nil {
			return err
		}

		tmpl, err := actionTmpl.Parse(string(contents))
		if err != nil {
			return err
		}
		contentTemplate[ext] = tmpl

	}

	return err
}

type xhrFileReader struct {
	rootUrl *url.URL
}

func (x xhrFileReader) ReadFile(filename string) ([]byte, error) {
	fmt.Printf("Read %v, %v\n", filename, x.rootUrl.String())
	v, err := url.Parse(filename)
	if err != nil {
		return nil, err
	}

	u := x.rootUrl.ResolveReference(v).String()

	req := xhr.NewRequest("GET", u)
	err = req.Send(nil)
	if err != nil {
		return nil, err
	}

	return []byte(req.ResponseText), nil
}

// renderDoc reads the present file, gets its template representation,
// and executes the template, sending output to w.
func renderDoc(w io.Writer, docUrl string, r io.Reader) error {
	u, err := url.Parse(docUrl)
	if err != nil {
		return err
	}

	xhrRead := xhrFileReader{
		rootUrl: u,
	}

	ctxt := &present.Context{
		ReadFile: xhrRead.ReadFile,
	}

	docFile := path.Base(docUrl)

	// Read the input and build the doc structure.
	doc, err := ctxt.Parse(r, docFile, 0)
	if err != nil {
		return err
	}

	for _, s := range doc.Sections {
		for i, e := range s.Elem {
			switch v := e.(type) {
			case present.Image:
				vv, err := url.Parse(v.URL)
				if err != nil {
					panic(err)
				}

				v.URL = u.ResolveReference(vv).String()
				s.Elem[i] = v
			}
		}
	}

	// Find which template should be executed.
	tmpl := contentTemplate[filepath.Ext(docFile)]

	mw := io.MultiWriter(w)

	// Execute the template.
	return doc.Render(mw, tmpl)
}

// showFile reports whether the given file should be displayed in the list.
func showFile(n string) bool {
	switch filepath.Ext(n) {
	case ".pdf":
	case ".html":
	case ".go":
	default:
		return isDoc(n)
	}
	return true
}

// showDir reports whether the given directory should be displayed in the list.
func showDir(n string) bool {
	if len(n) > 0 && (n[0] == '.' || n[0] == '_') || n == "present" {
		return false
	}
	return true
}

type dirListData struct {
	Path                          string
	Dirs, Slides, Articles, Other dirEntrySlice
}

type dirEntry struct {
	Name, Path, Title string
}

type dirEntrySlice []dirEntry

func (s dirEntrySlice) Len() int           { return len(s) }
func (s dirEntrySlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s dirEntrySlice) Less(i, j int) bool { return s[i].Name < s[j].Name }

func playable(c present.Code) bool {
	return present.PlayEnabled && c.Play && c.Ext == ".go"
}
