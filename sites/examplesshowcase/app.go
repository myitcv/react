package main

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples"
)

type tab uint32

const (
	tabShowcase tab = iota
	tabImmutable
)

type AppDef struct {
	r.ComponentDef
}

type AppState struct {
	tab
}

func App() *AppDef {
	res := new(AppDef)

	r.BlessElement(res, nil)

	return res
}

func (a *AppDef) Render() r.Element {
	var view r.Element

	switch a.State().tab {
	case tabShowcase:
		view = examples.Examples()
	case tabImmutable:
		view = examples.ImmExamples()
	}

	return r.Div(nil,
		r.Nav(
			r.NavProps(func(np *r.NavPropsDef) {
				np.ClassName = "navbar navbar-inverse navbar-fixed-top"
			}),
			r.Div(
				r.DivProps(func(dp *r.DivPropsDef) {
					dp.ClassName = "container"
				}),
				r.Div(
					r.DivProps(func(dp *r.DivPropsDef) {
						dp.ClassName = "navbar-header"
					}),
					r.A(
						r.AProps(func(ap *r.APropsDef) {
							ap.ClassName = "navbar-brand"
							ap.Href = "#"
						}),
						r.S("GopherJS React Examples"),
					),
				),
				r.Div(
					r.DivProps(func(dp *r.DivPropsDef) {
						dp.ID = "navbar"
						dp.ClassName = "collapse navbar-collapse"
					}),
					r.Ul(
						r.UlProps(func(ul *r.UlPropsDef) {
							ul.ClassName = "nav navbar-nav"
						}),
						a.buildLink("Showcase", tabShowcase, a.selectShowcase),
						a.buildLink("Immutable", tabImmutable, a.selectImmutable),
					),
				),
			),
		),
		r.Div(
			r.DivProps(func(dp *r.DivPropsDef) {
				dp.ClassName = "container"
			}),
			r.Div(
				r.DivProps(func(dp *r.DivPropsDef) {
					dp.ClassName = "starter-template"
				}),
				view,
			),
		),
	)
}

func (a *AppDef) buildLink(n string, t tab, cb func(e *r.SyntheticMouseEvent)) *r.LiDef {
	return r.Li(
		r.LiProps(func(li *r.LiPropsDef) {
			if a.State().tab == t {
				li.ClassName = "active"
			}
		}),
		r.A(
			r.AProps(func(ap *r.APropsDef) {
				ap.Href = "#"
				ap.OnClick = cb
			}),
			r.S(n),
		),
	)
}

func (a *AppDef) selectShowcase(e *r.SyntheticMouseEvent) {
	a.SetState(AppState{tabShowcase})
	e.PreventDefault()
}

func (a *AppDef) selectImmutable(e *r.SyntheticMouseEvent) {
	a.SetState(AppState{tabImmutable})
	e.PreventDefault()
}
