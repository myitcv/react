// Code generated by reactGen. DO NOT EDIT.

package html

import "myitcv.io/react/dom"

// H4Props defines the properties for the <h4> element
type H4Props struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	ID                      string
	Key                     string
	OnChange                dom.OnChange
	OnClick                 dom.OnClick
	Role                    string
	Style                   *CSS
}

func (h *H4Props) assign(v *_H4Props) {

	v.ClassName = h.ClassName

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	v.OnChange = h.OnChange

	v.OnClick = h.OnClick

	v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = h.Style.hack()

}