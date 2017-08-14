package todoapp // import "myitcv.io/react/examples/todoapp"

import (
	"fmt"

	"honnef.co/go/js/dom"
	"myitcv.io/react"
)

//go:generate reactGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	react.ComponentDef
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
func (t TodoAppDef) Render() react.Element {
	var entries []*react.LiElem

	for _, v := range t.State().items {
		entry := react.Li(nil, react.S(v))
		entries = append(entries, entry)
	}

	return react.Div(nil,
		react.H3(nil, react.S("TODO")),
		react.Ul(nil, entries...),
		react.Form(&react.FormProps{ClassName: "form-inline"},
			react.Div(
				&react.DivProps{ClassName: "form-group"},
				react.Label(&react.LabelProps{ClassName: "sr-only", For: "todoText"}, react.S("Todo Item")),
				react.Input(&react.InputProps{
					Type:        "text",
					ClassName:   "form-control",
					ID:          "todoText",
					Placeholder: "Todo Item",
					Value:       t.State().currItem,
					OnChange:    inputChange{t},
				}),
				react.Button(&react.ButtonProps{
					Type:      "submit",
					ClassName: "btn btn-default",
					OnClick:   add{t},
				}, react.S(fmt.Sprintf("Add #%v", len(t.State().items)+1))),
			),
		),
	)
}

type inputChange struct{ t TodoAppDef }
type add struct{ t TodoAppDef }

func (i inputChange) OnChange(se *react.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := i.t.State()
	ns.currItem = target.Value

	i.t.SetState(ns)
}

func (a add) OnClick(se *react.SyntheticMouseEvent) {
	ns := a.t.State()
	ns.items = append(ns.items, ns.currItem)
	ns.currItem = ""

	a.t.SetState(ns)

	se.PreventDefault()
}
