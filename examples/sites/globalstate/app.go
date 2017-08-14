package main

import (
	"myitcv.io/react"
	"myitcv.io/react/examples/sites/globalstate/state"
)

type AppDef struct {
	react.ComponentDef
}

type AppState struct {
	hideViewer bool
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() react.Element {
	var viewer *react.DivElem
	var showHide *react.ButtonElem

	if a.State().hideViewer {
		showHide = react.Button(
			&react.ButtonProps{OnClick: hideClick{a, false}},
			react.S("Show viewer"),
		)
	} else {
		viewer = react.Div(nil,
			react.H3(nil, react.S("Person Viewer")),
			PersonViewer(),
		)
		showHide = react.Button(
			&react.ButtonProps{OnClick: hideClick{a, true}},
			react.S("Hide viewer"),
		)
	}

	return react.Div(&react.DivProps{ClassName: "container"},
		react.H3(nil, react.S("Person Chooser")),
		PersonChooser(PersonChooserProps{
			PersonState: state.State.CurrentPerson(),
		}),
		react.Hr(nil),
		showHide,
		viewer,
	)
}

type hideClick struct {
	AppDef
	showHide bool
}

func (h hideClick) OnClick(e *react.SyntheticMouseEvent) {
	h.AppDef.SetState(AppState{
		hideViewer: h.showHide,
	})
}
