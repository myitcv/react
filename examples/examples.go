// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package examples // import "myitcv.io/react/examples"

import (
	"honnef.co/go/js/xhr"
	"myitcv.io/highlightjs"
	"myitcv.io/react"
	"myitcv.io/react/examples/hellomessage"
	"myitcv.io/react/examples/markdowneditor"
	"myitcv.io/react/examples/timer"
	"myitcv.io/react/examples/todoapp"
	"myitcv.io/react/jsx"
)

//go:generate reactGen
//go:generate immutableGen

// ExamplesDef is the definition of the Examples component
type ExamplesDef struct {
	react.ComponentDef
}

type tab int

const (
	tabGo tab = iota
	tabJsx
)

// Examples creates instances of the Examples component
func Examples() *ExamplesElem {
	return buildExamplesElem()
}

type (
	_Imm_tabS map[exampleKey]tab
)

// ExamplesState is the state type for the Examples component
type ExamplesState struct {
	examples     *exampleSource
	selectedTabs *tabS
}

// ComponentWillMount is a React lifecycle method for the Examples component
func (p ExamplesDef) ComponentWillMount() {
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

// GetInitialState returns in the initial state for the Examples component
func (p ExamplesDef) GetInitialState() ExamplesState {
	return ExamplesState{
		examples:     sources,
		selectedTabs: newTabS(),
	}
}

// Render renders the Examples component
func (p ExamplesDef) Render() react.Element {
	dc := jsx.HTML(`
		<h3>Introduction</h3>

		<p>This entire page is a React application. An outer <code>Examples</code> component
		contains a number of inner components.</p>

		<p>For the source code, raising issues, questions etc, please see
		<a href="https://github.com/myitcv/react/tree/master/examples" target="_blank">the Github repo</a>.</p>

		<p>Note the examples below show the Go source code from <code>master</code>.</p>
		`)

	dc = append(dc,
		p.renderExample(
			exampleHello,
			react.S("A Simple Example"),
			react.P(nil, react.S("The hellomessage.HelloMessage component demonstrates the simple use of a Props type.")),
			helloMessageJsx,
			hellomessage.HelloMessage(hellomessage.HelloMessageProps{Name: "Jane"}),
		),

		react.Hr(nil),

		p.renderExample(
			exampleTimer,
			react.S("A Stateful Component"),
			react.P(nil, react.S("The timer.Timer component demonstrates the use of a State type.")),
			timerJsx,
			timer.Timer(),
		),

		react.Hr(nil),

		p.renderExample(
			exampleTodo,
			react.S("An Application"),
			react.P(nil, react.S("The todoapp.TodoApp component demonstrates the use of state and event handling, but also the "+
				"problems of having a non-comparable state struct type.")),
			applicationJsx,
			todoapp.TodoApp(),
		),

		react.Hr(nil),

		p.renderExample(
			exampleMarkdown,
			react.S("A Component Using External Plugins"),
			react.P(nil, react.S("The markdowneditor.MarkdownEditor component demonstrates the use of an external Javascript library.")),
			markdownEditorJsx,
			markdowneditor.MarkdownEditor(),
		),

		react.Hr(nil),

		p.renderExample(
			exampleLatency,
			react.S("Latency Checker"),
			react.P(nil,
				react.S("By kind permission of "), react.A(&react.AProps{Href: "http://tjholowaychuk.com/"}, react.S("TJ Holowaychuk")),
				react.S(", a basic, component-based version of the beautiful APEX "), react.A(&react.AProps{Href: "https://latency.apex.sh/"}, react.S("Latency Tool")),
				react.S(" that uses randomly generated latency values."),
			),
			latencyJsx,
			react.A(&react.AProps{Href: "../latency", Target: "_blank"}, react.S("Launch in new tab")),
		),
	)

	return react.Div(&react.DivProps{ClassName: "container"},
		dc...,
	)
}

func (p ExamplesDef) renderExample(key exampleKey, title, msg react.Element, jsxSrc string, elem react.Element) react.Element {

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

func (p ExamplesDef) buildExampleNavTab(key exampleKey, t tab, title string) *react.LiElem {
	lip := &react.LiProps{Role: "presentation"}

	if v, _ := p.State().selectedTabs.Get(key); v == t {
		lip.ClassName = "active"
	}

	return react.Li(
		lip,
		react.A(
			&react.AProps{Href: "#", OnClick: tabChange{p, key, t}},
			react.S(title),
		),
	)

}

type tabChange struct {
	e   ExamplesDef
	key exampleKey
	t   tab
}

func (tc tabChange) OnClick(e *react.SyntheticMouseEvent) {
	p := tc.e
	key := tc.key
	t := tc.t

	cts := p.State().selectedTabs
	newSt := p.State()

	newSt.selectedTabs = cts.Set(key, t)
	p.SetState(newSt)

	e.PreventDefault()
}

func plainPanel(children ...react.Element) react.Element {
	return react.Div(&react.DivProps{ClassName: "panel panel-default panel-body"},
		children...,
	)
}
