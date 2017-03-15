// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package immtodoapp

import (
	"fmt"

	r "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/dom"
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
func TodoApp() *TodoAppDef {
	res := &TodoAppDef{}
	r.BlessElement(res, nil)
	return res
}

func (t *TodoAppDef) GetInitialState() TodoAppState {
	return TodoAppState{
		items: new(itemS),
	}
}

// Render renders the TodoApp component
func (t *TodoAppDef) Render() r.Element {
	var entries []*r.LiDef

	for _, v := range t.State().items.Range() {
		entry := r.Li(nil, r.S(v.name()))
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
						ip.ID = "todoText"
						ip.Placeholder = "Todo Item"
						ip.Value = t.State().currItem
						ip.OnChange = t.onCurrItemChange
					}),
				),
				r.Button(
					r.ButtonProps(func(bp *r.ButtonPropsDef) {
						bp.Type = "submit"
						bp.ClassName = "btn btn-default"
						bp.OnClick = t.onAddClicked
					}),
					r.S(fmt.Sprintf("Add #%v", t.State().items.Len()+1)),
				),
			),
		),
	)
}

func (t *TodoAppDef) onCurrItemChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := t.State()
	ns.currItem = target.Value

	t.SetState(ns)
}

func (t *TodoAppDef) onAddClicked(se *r.SyntheticMouseEvent) {
	ns := t.State()

	ns.items = ns.items.Append(new(item).setName(ns.currItem))
	ns.currItem = ""

	t.SetState(ns)

	se.PreventDefault()
}
