// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// DangerousInnerHTMLDef is convenience component definition that allows HTML to be directly
// set as the child of a DOM element. See
// https://facebook.github.io/react/docs/dom-elements.html#dangerouslysetinnerhtml for more details
type DangerousInnerHTMLDef struct {
	o *js.Object
}

// DangerousInnerHTML creates a new instance of a DangerousInnerHTMLDef component, using the
// supplied string as the raw HTML
func DangerousInnerHTML(s string) *DangerousInnerHTMLDef {
	o := object.New()
	o.Set("__html", s)

	res := &DangerousInnerHTMLDef{o: o}

	return res
}

func (d *DangerousInnerHTMLDef) reactElement() {}
