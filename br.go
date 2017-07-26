// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// BRElem is the React element definition corresponding to the HTML <br> element
type BRElem struct {
	Element
}

// _BRProps defines the properties for the <br> element
type _BRProps struct {
	*BasicHTMLElement
}

// BR creates a new instance of a <br> element with the provided props
func BR(props *BRProps) *BRElem {

	rProps := &_BRProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "br", rProps)

	return &BRElem{Element: elementHolder{elem: underlying}}
}
