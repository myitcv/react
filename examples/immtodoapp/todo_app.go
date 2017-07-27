package immtodoapp // import "myitcv.io/react/examples/immtodoapp"

import (
	"fmt"

	"honnef.co/go/js/dom"
	r "myitcv.io/react"
)

//go:generate reactGen
//go:generate immutableGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	r.ComponentDef
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
func (t TodoAppDef) Render() r.Element {
	var entries []*r.LiElem

	for _, v := range t.State().items.Range() {
		entry := r.Li(nil, r.S(v.name()))
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
				}, r.S(fmt.Sprintf("Add #%v", t.State().items.Len()+1))),
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

	ns.items = ns.items.Append(new(item).setName(ns.currItem))
	ns.currItem = ""

	a.t.SetState(ns)

	se.PreventDefault()
}
