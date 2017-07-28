// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package html

import (
	"github.com/gopherjs/gopherjs/js"
	"myitcv.io/react/dom"
)

type BasicNode struct {
	o *js.Object
}

type BasicElement struct {
	*BasicNode
}

func newBasicElement() *BasicElement {
	return &BasicElement{
		BasicNode: &BasicNode{object.New()},
	}
}

type BasicHTMLElement struct {
	*BasicElement

	ID        string `js:"id" react:"omitempty"`
	Key       string `js:"key" react:"omitempty"`
	ClassName string `js:"className"`
	Role      string `js:"role"`
	Style     *CSS   `js:"style"`

	OnChange dom.OnChange `js:"onChange"`
	OnClick  dom.OnClick  `js:"onClick"`

	DangerouslySetInnerHTML *DangerousInnerHTML `js:"dangerouslySetInnerHTML"`
}

func newBasicHTMLElement() *BasicHTMLElement {
	return &BasicHTMLElement{
		BasicElement: newBasicElement(),
	}
}
