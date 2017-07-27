package immtodoapp // import "myitcv.io/react/examples/immtodoapp"

import (
	"fmt"

	"myitcv.io/react/html"

	"honnef.co/go/js/dom"
)

//go:generate reactGen
//go:generate immutableGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	html.ComponentDef
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
func (t TodoAppDef) Render() html.Element {
	var entries []*html.LiElem

	for _, v := range t.State().items.Range() {
		entry := html.Li(nil, html.S(v.name()))
		entries = append(entries, entry)
	}

	return html.Div(nil,
		html.H3(nil, html.S("TODO")),
		html.Ul(nil, entries...),
		html.Form(&html.FormProps{ClassName: "form-inline"},
			html.Div(
				&html.DivProps{ClassName: "form-group"},
				html.Label(&html.LabelProps{ClassName: "sr-only", For: "todoText"}, html.S("Todo Item")),
				html.Input(&html.InputProps{
					Type:        "text",
					ClassName:   "form-control",
					ID:          "todoText",
					Placeholder: "Todo Item",
					Value:       t.State().currItem,
					OnChange:    inputChange{t},
				}),
				html.Button(&html.ButtonProps{
					Type:      "submit",
					ClassName: "btn btn-default",
					OnClick:   add{t},
				}, html.S(fmt.Sprintf("Add #%v", t.State().items.Len()+1))),
			),
		),
	)
}

type inputChange struct{ t TodoAppDef }
type add struct{ t TodoAppDef }

func (i inputChange) OnChange(se *html.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := i.t.State()
	ns.currItem = target.Value

	i.t.SetState(ns)
}

func (a add) OnClick(se *html.SyntheticMouseEvent) {
	ns := a.t.State()

	ns.items = ns.items.Append(new(item).setName(ns.currItem))
	ns.currItem = ""

	a.t.SetState(ns)

	se.PreventDefault()
}
