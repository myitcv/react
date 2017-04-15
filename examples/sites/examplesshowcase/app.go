package main

import (
	r "myitcv.io/react"
	"myitcv.io/react/examples"
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
		r.Nav(&r.NavProps{ClassName: "navbar navbar-inverse navbar-fixed-top"},
			r.Div(&r.DivProps{ClassName: "container"},
				r.Div(&r.DivProps{ClassName: "navbar-header"},
					r.A(&r.AProps{ClassName: "navbar-brand", Href: "#"},
						r.S("GopherJS React Examples"),
					),
				),
				r.Div(&r.DivProps{ClassName: "collapse navbar-collapse", ID: "navbar"},
					r.Ul(&r.UlProps{ClassName: "nav navbar-nav"},
						a.buildLink("Showcase", tabShowcase, a.selectShowcase),
						a.buildLink("Immutable", tabImmutable, a.selectImmutable),
					),
				),
			),
		),
		r.Div(&r.DivProps{ClassName: "container"},
			r.Div(&r.DivProps{ClassName: "starter-template"},
				view,
			),
		),
	)
}

func (a *AppDef) buildLink(n string, t tab, cb func(e *r.SyntheticMouseEvent)) *r.LiDef {
	var lip *r.LiProps

	if a.State().tab == t {
		lip = &r.LiProps{ClassName: "active"}
	}

	return r.Li(lip,
		r.A(&r.AProps{Href: "#", OnClick: cb}, r.S(n)),
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
