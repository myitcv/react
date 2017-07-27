package todoapp // import "myitcv.io/react/examples/todoapp"

import (
	"fmt"

	"honnef.co/go/js/dom"
	r "myitcv.io/react"
)

//go:generate reactGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	r.ComponentDef
}

// TodoAppState is the state type for the TodoApp component
type TodoAppState struct {
	items    []string
	currItem string
}

// TodoApp creates instances of the TodoApp component
func TodoApp() *TodoAppElem {
	return buildTodoAppElem()
}

// Equals must be defined because struct val instances of TodoAppState cannot
// be compared. It is generally bad practice to have mutable values in state in
// this way; myitcv.io/immutable seeks to help address this problem.
// See myitcv.io/react/examples/immtodoapp for an example
func (c TodoAppState) Equals(v TodoAppState) bool {
	if c.currItem != v.currItem {
		return false
	}

	if len(v.items) != len(c.items) {
		return false
	}

	for i := range v.items {
		if v.items[i] != c.items[i] {
			return false
		}
	}

	return true
}

// Render renders the TodoApp component
func (t TodoAppDef) Render() r.Element {
	var entries []*r.LiElem

	for _, v := range t.State().items {
		entry := r.Li(nil, r.S(v))
		entries = append(entries, entry)
	}

	return r.Div(nil,
		r.H3(nil, r.S("TODO")),
		r.Ul(nil, entries...),
		r.Form(&r.FormProps{ClassName: "form-inline"},
			r.Div(
				&r.DivProps{ClassName: "form-group"},
				r.Label(&r.LabelProps{ClassName: "sr-only", For: "todoText"}, r.S("Todo Item")),
				r.Input(&r.InputProps{
					Type:        "text",
					ClassName:   "form-control",
					ID:          "todoText",
					Placeholder: "Todo Item",
					Value:       t.State().currItem,
					OnChange:    inputChange{t},
				}),
				r.Button(&r.ButtonProps{
					Type:      "submit",
					ClassName: "btn btn-default",
					OnClick:   add{t},
				}, r.S(fmt.Sprintf("Add #%v", len(t.State().items)+1))),
			),
		),
	)
}

type inputChange struct{ t TodoAppDef }
type add struct{ t TodoAppDef }

func (i inputChange) OnChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := i.t.State()
	ns.currItem = target.Value

	i.t.SetState(ns)
}

func (a add) OnClick(se *r.SyntheticMouseEvent) {
	ns := a.t.State()
	ns.items = append(ns.items, ns.currItem)
	ns.currItem = ""

	a.t.SetState(ns)

	se.PreventDefault()
}
