// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

// Package remarkable provides an incomplete wrapper for remarkable (https://github.com/jonschlinkert/remarkable),
// a pure Javascript markdown parser
//
package remarkable // import "myitcv.io/remarkable"

import "github.com/gopherjs/gopherjs/js"

var remarkable = js.Global.Get("Remarkable")

type Remarkable struct {
	o *js.Object

	// Render a markdown string s as an HTML string
	Render func(s string) string `js:"render"`
}

// NewRemarkable returns a new instance of remarkable
func NewRemarkable() *Remarkable {
	remark := remarkable.New()

	return &Remarkable{o: remark}
}
