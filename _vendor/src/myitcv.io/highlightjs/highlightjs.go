// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

// Package highlightjs provides an incomplete wrapper for Highlight.js (https://github.com/isagalaev/highlight.js),
// a Javascript syntax highlighter
//
package highlightjs // import "myitcv.io/highlightjs"

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func init() {
	// TODO this is a bit gross... and it's not configurable by language
	// but does avoid us having to add <link>'s to the <head> block

	doc := dom.GetWindow().Document()
	style := doc.CreateElement("style")
	style.SetTextContent(defaultCss)

	doc.GetElementsByTagName("head")[0].AppendChild(style)
}

var highlightjs = js.Global.Get("hljs")

// Highlight is the core highlighting function. See documentation at
// http://highlightjs.readthedocs.io/en/latest/api.html#highlight-name-value-ignore-illegals-continuation
//
func Highlight(lang string, source string, ignoreIllegals bool) *HighlightResult {
	o := highlightjs.Call("highlight", lang, source, ignoreIllegals)

	return &HighlightResult{o: o}
}

// HighlightResult represents the result of a highlighting
type HighlightResult struct {
	o *js.Object

	Language  string `js:"language"`
	Relevance int    `js:"relevance"`
	Value     string `js:"value"`
}
