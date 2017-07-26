// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// HRElem is the React element definition corresponding to the HTML <hr> element
type HRElem struct {
	Element
}

// _HRProps defines the properties for the <hr> element
type _HRProps struct {
	*BasicHTMLElement
}

// HR creates a new instance of a <hr> element with the provided props
func HR(props *HRProps) *HRElem {

	rProps := &_HRProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "hr", rProps)

	return &HRElem{Element: elementHolder{elem: underlying}}
}
