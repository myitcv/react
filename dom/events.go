package dom

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Event interface{}

type OnChange interface {
	Event

	OnChange(e *SyntheticEvent)
}

type OnClick interface {
	Event

	OnClick(e *SyntheticMouseEvent)
}

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
