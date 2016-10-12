package todoapp

import (
	"fmt"

	. "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/dom"
)

type TodoAppDef struct {
	ComponentDef
}

type TodoAppState struct {
	items    []string
	currItem string
}

func TodoApp() *TodoAppDef {
	res := &TodoAppDef{}

	BlessElement(res, nil)

	return res
}

func (p *TodoAppDef) GetInitialState() TodoAppState {
	return TodoAppState{}
}

func (p *TodoAppDef) Render() Element {
	var entries []*LiDef

	for _, v := range p.State().items {
		entry := Li(nil, S(v))
		entries = append(entries, entry)
	}

	// TODO why does this fail when inline below?
	theDp := DivProps(func(dp *DivPropsDef) {
		dp.ClassName = "form-group"
	})

	return Div(nil,
		H3(nil, S("TODO")),
		Ul(nil, entries...),
		Form(
			FormProps(func(fp *FormPropsDef) {
				fp.ClassName = "form-inline"
			}),
			Div(
				theDp,
				Label(
					LabelProps(func(lp *LabelPropsDef) {
						lp.ClassName = "sr-only"
						lp.For = "todoText"
					}),
					S("Todo Item"),
				),
				Input(
					InputProps(func(ip *InputPropsDef) {
						ip.Type = "text"
						ip.ClassName = "form-control"
						ip.Id = "todoText"
						ip.Placeholder = "Todo Item"
						ip.Value = p.State().currItem
						ip.OnChange = p.onCurrItemChange
					}),
				),
				Button(
					ButtonProps(func(bp *ButtonPropsDef) {
						bp.Type = "submit"
						bp.ClassName = "btn btn-default"
						bp.OnClick = p.onAddClicked
					}),
					S(fmt.Sprintf("Add #%v", len(p.State().items)+1)),
				),
			),
		),
	)
}

func (p *TodoAppDef) onCurrItemChange(se *SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := p.State()
	ns.currItem = target.Value

	p.SetState(ns)
}

func (p *TodoAppDef) onAddClicked(se *SyntheticMouseEvent) {
	se.PreventDefault()

	ns := p.State()
	ns.items = append(ns.items, ns.currItem)
	ns.currItem = ""

	p.SetState(ns)
}
