package examples

import (
	"honnef.co/go/js/xhr"
	"myitcv.io/highlightjs"
	"myitcv.io/react"
	"myitcv.io/react/examples/immtodoapp"
	"myitcv.io/react/jsx"
)

// ImmExamplesDef is the definition of the ImmExamples component
type ImmExamplesDef struct {
	react.ComponentDef
}

// ImmExamples creates instances of the ImmExamples component
func ImmExamples() *ImmExamplesElem {
	return buildImmExamplesElem()
}

// ImmExamplesState is the state type for the ImmExamples component
type ImmExamplesState struct {
	examples     *exampleSource
	selectedTabs *tabS
}

// ComponentWillMount is a React lifecycle method for the ImmExamples component
func (p ImmExamplesDef) ComponentWillMount() {
	if !fetchStarted {
		for i, e := range sources.Range() {
			go func(i exampleKey, e *source) {
				req := xhr.NewRequest("GET", "https://raw.githubusercontent.com/myitcv/react/master/examples/"+e.file())
				err := req.Send(nil)
				if err != nil {
					panic(err)
				}

				sources = sources.Set(i, e.setSrc(req.ResponseText))

				newSt := p.State()
				newSt.examples = sources
				p.SetState(newSt)
			}(i, e)
		}

		fetchStarted = true
	}
}

// GetInitialState returns in the initial state for the ImmExamples component
func (p ImmExamplesDef) GetInitialState() ImmExamplesState {
	return ImmExamplesState{
		examples:     sources,
		selectedTabs: newTabS(),
	}
}

// Render renders the ImmExamples component
func (p ImmExamplesDef) Render() react.Element {
	dc := jsx.HTML(`
		<h3>Using immutable data structures</h3>

		<p>This page focuses on using <a href="https://myitcv.io/immutable"><code>myitcv.io/immutable</code></a>
		(specifically <a href="https://github.com/myitcv/immutable/wiki/immutableGen"><code>immutableGen</code></a>) to
		help make building components easier. The pattern of immutable data structures lends itself well to React's style
		of composition.</p>

		<p>For the source code, raising issues, questions etc, please see
		<a href="https://github.com/myitcv/react/tree/master/examples" target="_blank">the Github repo</a>.</p>

		<p>Note the examples below show the Go source code from <code>master</code>.</p>
		`)

	dc = append(dc,
		p.renderExample(
			exampleImmTodo,
			react.Span(nil, react.S("A simple TODO app")),
			react.P(nil, react.S("The immtodoapp.TodoApp component is a reimplementation of todoapp.TodoApp using immutable data structures.")),
			"n/a",
			immtodoapp.TodoApp(),
		),
	)

	return react.Div(&react.DivProps{ClassName: "container"},
		dc...,
	)
}

func (p ImmExamplesDef) renderExample(key exampleKey, title, msg react.Element, jsxSrc string, elem react.Element) react.Element {

	var goSrc string
	src, _ := p.State().examples.Get(key)
	if src != nil {
		goSrc = src.src()
	}

	var code *react.DangerousInnerHTML
	switch v, _ := p.State().selectedTabs.Get(key); v {
	case tabGo:
		code = react.NewDangerousInnerHTML(highlightjs.Highlight("go", goSrc, true).Value)
	case tabJsx:
		code = react.NewDangerousInnerHTML(highlightjs.Highlight("javascript", jsxSrc, true).Value)
	}

	return react.Div(nil,
		react.H3(nil, title),
		msg,
		react.Div(&react.DivProps{ClassName: "row"},
			react.Div(&react.DivProps{ClassName: "col-md-8"},
				react.Div(&react.DivProps{ClassName: "panel panel-default with-nav-tabs"},
					react.Div(&react.DivProps{ClassName: "panel-heading"},
						react.Ul(
							&react.UlProps{ClassName: "nav nav-tabs"},

							p.buildExampleNavTab(key, tabGo, "GopherJS"),
							p.buildExampleNavTab(key, tabJsx, "JSX"),
						),
					),
					react.Div(&react.DivProps{ClassName: "panel-body"},
						react.Pre(&react.PreProps{
							Style: &react.CSS{
								MaxHeight: "400px",
							},
							DangerouslySetInnerHTML: code,
						}),
					),
				),
			),
			react.Div(&react.DivProps{ClassName: "col-md-4"},
				plainPanel(elem),
			),
		),
	)
}

func (p ImmExamplesDef) buildExampleNavTab(key exampleKey, t tab, title string) *react.LiElem {
	lip := &react.LiProps{Role: "presentation"}

	if v, _ := p.State().selectedTabs.Get(key); v == t {
		lip.ClassName = "active"
	}

	return react.Li(
		lip,
		react.A(
			&react.AProps{Href: "#", OnClick: immTabChange{p, key, t}},
			react.S(title),
		),
	)

}

type immTabChange struct {
	e   ImmExamplesDef
	key exampleKey
	t   tab
}

func (tc immTabChange) OnClick(e *react.SyntheticMouseEvent) {
	p := tc.e
	key := tc.key
	t := tc.t

	cts := p.State().selectedTabs
	newSt := p.State()

	newSt.selectedTabs = cts.Set(key, t)
	p.SetState(newSt)

	e.PreventDefault()
}

func (p *ImmExamplesDef) handleTabChange(key exampleKey, t tab) func(*react.SyntheticMouseEvent) {
	return func(e *react.SyntheticMouseEvent) {
		cts := p.State().selectedTabs
		newSt := p.State()

		newSt.selectedTabs = cts.Set(key, t)
		p.SetState(newSt)

		e.PreventDefault()
	}
}
