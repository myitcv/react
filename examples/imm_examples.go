package examples

import (
	"honnef.co/go/js/xhr"
	"myitcv.io/highlightjs"
	r "myitcv.io/react"
	"myitcv.io/react/examples/immtodoapp"
	"myitcv.io/react/jsx"
)

// ImmExamplesDef is the definition of the ImmExamples component
type ImmExamplesDef struct {
	r.ComponentDef
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
func (p ImmExamplesDef) Render() r.Element {
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
			r.Span(nil, r.S("A simple TODO app")),
			r.P(nil, r.S("The immtodoapp.TodoApp component is a reimplementation of todoapp.TodoApp using immutable data structures.")),
			"n/a",
			immtodoapp.TodoApp(),
		),
	)

	return r.Div(&r.DivProps{ClassName: "container"},
		dc...,
	)
}

func (p ImmExamplesDef) renderExample(key exampleKey, title, msg r.Element, jsxSrc string, elem r.Element) r.Element {

	var goSrc string
	src, _ := p.State().examples.Get(key)
	if src != nil {
		goSrc = src.src()
	}

	var code *r.DangerousInnerHTML
	switch v, _ := p.State().selectedTabs.Get(key); v {
	case tabGo:
		code = r.NewDangerousInnerHTML(highlightjs.Highlight("go", goSrc, true).Value)
	case tabJsx:
		code = r.NewDangerousInnerHTML(highlightjs.Highlight("javascript", jsxSrc, true).Value)
	}

	return r.Div(nil,
		r.H3(nil, title),
		msg,
		r.Div(&r.DivProps{ClassName: "row"},
			r.Div(&r.DivProps{ClassName: "col-md-8"},
				r.Div(&r.DivProps{ClassName: "panel panel-default with-nav-tabs"},
					r.Div(&r.DivProps{ClassName: "panel-heading"},
						r.Ul(
							&r.UlProps{ClassName: "nav nav-tabs"},

							p.buildExampleNavTab(key, tabGo, "GopherJS"),
							p.buildExampleNavTab(key, tabJsx, "JSX"),
						),
					),
					r.Div(&r.DivProps{ClassName: "panel-body"},
						r.Pre(&r.PreProps{
							Style: &r.CSS{
								MaxHeight: "400px",
							},
							DangerouslySetInnerHTML: code,
						}),
					),
				),
			),
			r.Div(&r.DivProps{ClassName: "col-md-4"},
				plainPanel(elem),
			),
		),
	)
}

func (p ImmExamplesDef) buildExampleNavTab(key exampleKey, t tab, title string) *r.LiElem {
	lip := &r.LiProps{Role: "presentation"}

	if v, _ := p.State().selectedTabs.Get(key); v == t {
		lip.ClassName = "active"
	}

	return r.Li(
		lip,
		r.A(
			&r.AProps{Href: "#", OnClick: immTabChange{p, key, t}},
			r.S(title),
		),
	)

}

type immTabChange struct {
	e   ImmExamplesDef
	key exampleKey
	t   tab
}

func (tc immTabChange) OnClick(e *r.SyntheticMouseEvent) {
	p := tc.e
	key := tc.key
	t := tc.t

	cts := p.State().selectedTabs
	newSt := p.State()

	newSt.selectedTabs = cts.Set(key, t)
	p.SetState(newSt)

	e.PreventDefault()
}

func (p *ImmExamplesDef) handleTabChange(key exampleKey, t tab) func(*r.SyntheticMouseEvent) {
	return func(e *r.SyntheticMouseEvent) {
		cts := p.State().selectedTabs
		newSt := p.State()

		newSt.selectedTabs = cts.Set(key, t)
		p.SetState(newSt)

		e.PreventDefault()
	}
}
