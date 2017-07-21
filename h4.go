// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// H4Def is the React component definition corresponding to the HTML <h4> element
type H4Def struct {
	underlying *js.Object
}

// _H4Props defines the properties for the <h4> element
type _H4Props struct {
	*BasicHTMLElement
}

func (d *H4Def) reactElement() {}

// H4 creates a new instance of a <h4> element with the provided props and
// child
func H4(props *H4Props, children ...Element) *H4Def {

	rProps := &_H4Props{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"h4", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &H4Def{underlying: underlying}
}
