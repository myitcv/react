// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import "github.com/gopherjs/gopherjs/js"

// ImgDef is the React component definition corresponding to the HTML <Img> element
type ImgDef struct {
	underlying *js.Object
}

// _ImgProps are the props for a <Img> component
type _ImgProps struct {
	*BasicHTMLElement

	Src string `js:"src"`
}

func (d *ImgDef) reactElement() {}

// Img creates a new instance of a <Img> element with the provided props and children
func Img(props *ImgProps, children ...Element) *ImgDef {

	rProps := &_ImgProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"Img", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &ImgDef{underlying: underlying}
}
