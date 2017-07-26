// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// AElem is the React element definition corresponding to the HTML <a> element
type AElem struct {
	Element
}

// _APropsDef defines the properties for the <a> element
type _AProps struct {
	*BasicHTMLElement

	Target string `js:"target"`
	Href   string `js:"href"`
}

// A creates a new instance of a <a> element with the provided props and
// children
func A(props *AProps, children ...Element) *AElem {

	rProps := &_AProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"a", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &AElem{Element: elementHolder{elem: underlying}}
}
