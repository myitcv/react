// Code generated by reactGen. DO NOT EDIT.

package html

import "myitcv.io/react/dom"

// BrProps defines the properties for the <br> element
type BrProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	ID                      string
	Key                     string
	OnChange                dom.OnChange
	OnClick                 dom.OnClick
	Role                    string
	Style                   *CSS
}

func (b *BrProps) assign(v *_BrProps) {

	v.ClassName = b.ClassName

	v.DangerouslySetInnerHTML = b.DangerouslySetInnerHTML

	if b.ID != "" {
		v.ID = b.ID
	}

	if b.Key != "" {
		v.Key = b.Key
	}

	v.OnChange = b.OnChange

	v.OnClick = b.OnClick

	v.Role = b.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = b.Style.hack()

}