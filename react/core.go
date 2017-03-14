package react

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
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

	ID        string `js:"id"`
	Key       string `js:"key"`
	ClassName string `js:"className"`

	OnChange func(e *SyntheticEvent)      `js:"onChange"`
	OnClick  func(e *SyntheticMouseEvent) `js:"onClick"`

	DangerouslySetInnerHTML *DangerousInnerHTMLDef `js:"dangerouslySetInnerHTML"`
}

func newBasicHTMLElement() *BasicHTMLElement {
	return &BasicHTMLElement{
		BasicElement: newBasicElement(),
	}
}

// TODO complete the definition
type SyntheticEvent struct {
	o *js.Object

	PreventDefault  func() `js:"preventDefault"`
	StopPropagation func() `js:"stopPropagation"`
}

func (s *SyntheticEvent) Target() dom.HTMLElement {
	return dom.WrapHTMLElement(s.o.Get("target"))
}

type SyntheticMouseEvent struct {
	*SyntheticEvent

	ClientX int `js:"clientX"`
}
