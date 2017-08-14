package main

import (
	"myitcv.io/react"
	"myitcv.io/react/examples"
)

type tab uint32

const (
	tabShowcase tab = iota
	tabImmutable
	tabGlobalState
)

type AppDef struct {
	react.ComponentDef
}

type AppState struct {
	tab
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() react.Element {
	var view react.Element

	switch a.State().tab {
	case tabShowcase:
		view = examples.Examples()
	case tabImmutable:
		view = examples.ImmExamples()
	case tabGlobalState:
		view = examples.GlobalStateExamples()
	}

	return react.Div(nil,
		react.Nav(&react.NavProps{ClassName: "navbar navbar-inverse navbar-fixed-top"},
			react.Div(&react.DivProps{ClassName: "container"},
				react.Div(&react.DivProps{ClassName: "navbar-header"},
					react.A(&react.AProps{ClassName: "navbar-brand", Href: "#"},
						react.S("GopherJS React Examples"),
					),
				),
				react.Div(&react.DivProps{ClassName: "collapse navbar-collapse", ID: "navbar"},
					react.Ul(&react.UlProps{ClassName: "nav navbar-nav"},
						a.buildLink("Simple", tabShowcase, tabChange{a, tabShowcase}),
						a.buildLink("Immutable", tabImmutable, tabChange{a, tabImmutable}),
						a.buildLink("Global State", tabGlobalState, tabChange{a, tabGlobalState}),
					),
				),
			),
		),
		react.Div(&react.DivProps{ClassName: "container"},
			react.Div(&react.DivProps{ClassName: "starter-template"},
				view,
			),
		),
	)
}

type tabChange struct {
	a AppDef
	t tab
}

func (tc tabChange) OnClick(e *react.SyntheticMouseEvent) {
	tc.a.SetState(AppState{tc.t})
	e.PreventDefault()
}

func (a AppDef) buildLink(n string, t tab, tc tabChange) *react.LiElem {
	var lip *react.LiProps

	if a.State().tab == t {
		lip = &react.LiProps{ClassName: "active"}
	}

	return react.Li(lip,
		react.A(&react.AProps{Href: "#", OnClick: tc}, react.S(n)),
	)
}
