package react

import "github.com/gopherjs/gopherjs/js"

type Event interface{}

type Ref interface {
	Ref(h *js.Object)
}

type OnChange interface {
	Event

	OnChange(e *SyntheticEvent)
}

type OnClick interface {
	Event

	OnClick(e *SyntheticMouseEvent)
}
