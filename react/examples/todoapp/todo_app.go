// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package todoapp

import (
	"fmt"

	r "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/dom"
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
func TodoApp() *TodoAppDef {
	res := &TodoAppDef{}
	r.BlessElement(res, nil)
	return res
}

// Equals must be defined because struct val instances of TodoAppState cannot
// be compared. It is generally bad practice to have mutable values in state in
// this way; github.com/myitcv/immutable seeks to help address this problem.
// See github.com/myitcv/gopherjs/react/examples/immtodoapp for an example
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
func (t *TodoAppDef) Render() r.Element {
	var entries []*r.LiDef

	for _, v := range t.State().items {
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
					r.S(fmt.Sprintf("Add #%v", len(t.State().items)+1)),
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
	ns.items = append(ns.items, ns.currItem)
	ns.currItem = ""

	t.SetState(ns)

	se.PreventDefault()
}
