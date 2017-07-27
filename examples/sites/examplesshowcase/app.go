package main

import (
	r "myitcv.io/react"
	"myitcv.io/react/examples"
)

type tab uint32

const (
	tabShowcase tab = iota
	tabImmutable
	tabGlobalState
)

type AppDef struct {
	r.ComponentDef
}

type AppState struct {
	tab
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() r.Element {
	var view r.Element

	switch a.State().tab {
	case tabShowcase:
		view = examples.Examples()
	case tabImmutable:
		view = examples.ImmExamples()
	case tabGlobalState:
		view = examples.GlobalStateExamples()
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
						a.buildLink("Simple", tabShowcase, tabChange{a, tabShowcase}),
						a.buildLink("Immutable", tabImmutable, tabChange{a, tabImmutable}),
						a.buildLink("Global State", tabGlobalState, tabChange{a, tabGlobalState}),
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

type tabChange struct {
	a AppDef
	t tab
}

func (tc tabChange) OnClick(e *r.SyntheticMouseEvent) {
	tc.a.SetState(AppState{tc.t})
	e.PreventDefault()
}

func (a AppDef) buildLink(n string, t tab, tc tabChange) *r.LiElem {
	var lip *r.LiProps

	if a.State().tab == t {
		lip = &r.LiProps{ClassName: "active"}
	}

	return r.Li(lip,
		r.A(&r.AProps{Href: "#", OnClick: tc}, r.S(n)),
	)
}
