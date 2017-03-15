// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// H3Def is the React component definition corresponding to the HTML <h3> element
type H3Def struct {
	underlying *js.Object
}

// H3PropsDef defines the properties for the <h3> element
type H3PropsDef struct {
	*BasicHTMLElement
}

// H3Props creates a new instance of <h3> properties, mutating the props
// by the provided initiah3ser
func H3Props(f func(p *H3PropsDef)) *H3PropsDef {
	res := &H3PropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *H3Def) reactElement() {}

// H3 creates a new instance of a <h3> element with the provided props
func H3(props *H3PropsDef, child Element) *H3Def {
	underlying := react.Call("createElement", "h3", props, elementToReactObj(child))

	return &H3Def{underlying: underlying}
}
