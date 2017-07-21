// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// FooterDef is the React component definition corresponding to the HTML <footer> element
type FooterDef struct {
	underlying *js.Object
}

// _FooterProps are the props for a <footer> component
type _FooterProps struct {
	*BasicHTMLElement
}

func (d *FooterDef) reactElement() {}

// Footer creates a new instance of a <footer> element with the provided props and children
func Footer(props *FooterProps, children ...Element) *FooterDef {

	rProps := &_FooterProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"footer", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &FooterDef{underlying: underlying}
}
