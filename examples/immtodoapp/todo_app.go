package immtodoapp // import "myitcv.io/react/examples/immtodoapp"

import (
	"fmt"

	"honnef.co/go/js/dom"
	"myitcv.io/react"
)

//go:generate reactGen
//go:generate immutableGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	react.ComponentDef
}

type _Imm_item struct {
	name string
}

type _Imm_itemS []*item

// TodoAppState is the state type for the TodoApp component
type TodoAppState struct {
	items    *itemS
	currItem string
}

// TodoApp creates instances of the TodoApp component
func TodoApp() *TodoAppElem {
	return buildTodoAppElem()
}

func (t TodoAppDef) GetInitialState() TodoAppState {
	return TodoAppState{
		items: new(itemS),
	}
}

// Render renders the TodoApp component
func (t TodoAppDef) Render() react.Element {
	var entries []*react.LiElem

	for _, v := range t.State().items.Range() {
		entry := react.Li(nil, react.S(v.name()))
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
				}, react.S(fmt.Sprintf("Add #%v", t.State().items.Len()+1))),
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

	ns.items = ns.items.Append(new(item).setName(ns.currItem))
	ns.currItem = ""

	a.t.SetState(ns)

	se.PreventDefault()
}
