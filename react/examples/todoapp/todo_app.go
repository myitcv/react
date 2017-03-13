package todoapp

import (
	"fmt"

	r "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/dom"
)

type TodoAppDef struct {
	r.ComponentDef
}

type TodoAppState struct {
	items    []string
	currItem string
}

func TodoApp() *TodoAppDef {
	res := &TodoAppDef{}

	r.BlessElement(res, nil)

	return res
}

func (p *TodoAppDef) GetInitialState() TodoAppState {
	return TodoAppState{}
}

func (p *TodoAppDef) Render() r.Element {
	var entries []*r.LiDef

	for _, v := range p.State().items {
		entry := r.Li(nil, r.S(v))
		entries = append(entries, entry)
	}

	// TODO why does this fail when inline below?
	theDp := r.DivProps(func(dp *r.DivPropsDef) {
		dp.ClassName = "form-group"
	})

	return r.Div(nil,
		r.H3(nil, r.S("TODO")),
		r.Ul(nil, entries...),
		r.Form(
			r.FormProps(func(fp *r.FormPropsDef) {
				fp.ClassName = "form-inline"
			}),
			r.Div(
				theDp,
				r.Label(
					r.LabelProps(func(lp *r.LabelPropsDef) {
						lp.ClassName = "sr-only"
						lp.For = "todoText"
					}),
					r.S("Todo Item"),
				),
				r.Input(
					r.InputProps(func(ip *r.InputPropsDef) {
						ip.Type = "text"
						ip.ClassName = "form-control"
						ip.Id = "todoText"
						ip.Placeholder = "Todo Item"
						ip.Value = p.State().currItem
						ip.OnChange = p.onCurrItemChange
					}),
				),
				r.Button(
					r.ButtonProps(func(bp *r.ButtonPropsDef) {
						bp.Type = "submit"
						bp.ClassName = "btn btn-default"
						bp.OnClick = p.onAddClicked
					}),
					r.S(fmt.Sprintf("Add #%v", len(p.State().items)+1)),
				),
			),
		),
	)
}

func (p *TodoAppDef) onCurrItemChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := p.State()
	ns.currItem = target.Value

	p.SetState(ns)
}

func (p *TodoAppDef) onAddClicked(se *r.SyntheticMouseEvent) {
	se.PreventDefault()

	ns := p.State()
	ns.items = append(ns.items, ns.currItem)
	ns.currItem = ""

	p.SetState(ns)

	se.PreventDefault()
}
