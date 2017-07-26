// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// PreElem is the React element definition corresponding to the HTML <pre> element
type PreElem struct {
	Element
}

// _PreProps defines the properties for the <pre> element
type _PreProps struct {
	*BasicHTMLElement
}

// Pre creates a new instance of a <pre> element with the provided props and
// children
func Pre(props *PreProps, children ...Element) *PreElem {

	rProps := &_PreProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"pre", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &PreElem{Element: elementHolder{elem: underlying}}
}
