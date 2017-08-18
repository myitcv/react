// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io"
	"net/url"
	"path"
	"path/filepath"

	"honnef.co/go/js/xhr"

	"golang.org/x/tools/present"
)

// func init() {
// 	http.HandleFunc("/", dirHandler)
// }

// func dirHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/favicon.ico" {
// 		http.Error(w, "not found", 404)
// 		return
// 	}
// 	const base = "."
// 	name := filepath.Join(base, r.URL.Path)
// 	if isDoc(name) {
// 		err := renderDoc(w, name)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, err.Error(), 500)
// 		}
// 		return
// 	}
// 	if isDir, err := dirList(w, name); err != nil {
// 		addr, _, e := net.SplitHostPort(r.RemoteAddr)
// 		if e != nil {
// 			addr = r.RemoteAddr
// 		}
// 		log.Printf("request from %s: %s", addr, err)
// 		http.Error(w, err.Error(), 500)
// 		return
// 	} else if isDir {
// 		return
// 	}
// 	http.FileServer(http.Dir(base)).ServeHTTP(w, r)
// }

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

		// Read and parse the input.
		// tmpl := present.Template()
		// tmpl = tmpl.Funcs(template.FuncMap{"playable": playable})
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
	rootUrl string
}

func (x xhrFileReader) ReadFile(filename string) ([]byte, error) {
	url := x.rootUrl + "/" + filename

	req := xhr.NewRequest("GET", url)
	err := req.Send(nil)
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

	b := *u
	b.Path = path.Dir(b.Path)

	xhrRead := xhrFileReader{
		rootUrl: b.String(),
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
				// u, err := url.Parse(v.URL)
				// if err != nil {
				// 	panic(err)
				// }

				v.URL = b.String() + "/" + v.URL
				s.Elem[i] = v

				// v.URL = b.String() + "/" + v.UR

			}
		}
	}

	// Find which template should be executed.
	tmpl := contentTemplate[filepath.Ext(docFile)]

	mw := io.MultiWriter(w)

	// Execute the template.
	return doc.Render(mw, tmpl)
}

// func parse(name io.Reader, mode present.ParseMode) (*present.Doc, error) {
// 	return present.Parse(name, name, 0)
// }

// dirList scans the given path and writes a directory listing to w.
// It parses the first part of each .slide file it encounters to display the
// presentation title in the listing.
// If the given path is not a directory, it returns (isDir == false, err == nil)
// and writes nothing to w.
// func dirList(w io.Writer, name string) (isDir bool, err error) {
// 	f, err := os.Open(name)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer f.Close()
// 	fi, err := f.Stat()
// 	if err != nil {
// 		return false, err
// 	}
// 	if isDir = fi.IsDir(); !isDir {
// 		return false, nil
// 	}
// 	fis, err := f.Readdir(0)
// 	if err != nil {
// 		return false, err
// 	}
// 	d := &dirListData{Path: name}
// 	for _, fi := range fis {
// 		// skip the golang.org directory
// 		if name == "." && fi.Name() == "golang.org" {
// 			continue
// 		}
// 		e := dirEntry{
// 			Name: fi.Name(),
// 			Path: filepath.ToSlash(filepath.Join(name, fi.Name())),
// 		}
// 		if fi.IsDir() && showDir(e.Name) {
// 			d.Dirs = append(d.Dirs, e)
// 			continue
// 		}
// 		if isDoc(e.Name) {
// 			if p, err := parse(e.Path, present.TitlesOnly); err != nil {
// 				log.Println(err)
// 			} else {
// 				e.Title = p.Title
// 			}
// 			switch filepath.Ext(e.Path) {
// 			case ".article":
// 				d.Articles = append(d.Articles, e)
// 			case ".slide":
// 				d.Slides = append(d.Slides, e)
// 			}
// 		} else if showFile(e.Name) {
// 			d.Other = append(d.Other, e)
// 		}
// 	}
// 	if d.Path == "." {
// 		d.Path = ""
// 	}
// 	sort.Sort(d.Dirs)
// 	sort.Sort(d.Slides)
// 	sort.Sort(d.Articles)
// 	sort.Sort(d.Other)
// 	return true, dirListTemplate.Execute(w, d)
// }

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
