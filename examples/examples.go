// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package examples // import "myitcv.io/react/examples"

import (
	"honnef.co/go/js/xhr"
	"myitcv.io/highlightjs"
	r "myitcv.io/react"
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
	r.ComponentDef
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
func (p ExamplesDef) Render() r.Element {
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
			r.S("A Simple Example"),
			r.P(nil, r.S("The hellomessage.HelloMessage component demonstrates the simple use of a Props type.")),
			helloMessageJsx,
			hellomessage.HelloMessage(hellomessage.HelloMessageProps{Name: "Jane"}),
		),

		r.Hr(nil),

		p.renderExample(
			exampleTimer,
			r.S("A Stateful Component"),
			r.P(nil, r.S("The timer.Timer component demonstrates the use of a State type.")),
			timerJsx,
			timer.Timer(),
		),

		r.Hr(nil),

		p.renderExample(
			exampleTodo,
			r.S("An Application"),
			r.P(nil, r.S("The todoapp.TodoApp component demonstrates the use of state and event handling, but also the "+
				"problems of having a non-comparable state struct type.")),
			applicationJsx,
			todoapp.TodoApp(),
		),

		r.Hr(nil),

		p.renderExample(
			exampleMarkdown,
			r.S("A Component Using External Plugins"),
			r.P(nil, r.S("The markdowneditor.MarkdownEditor component demonstrates the use of an external Javascript library.")),
			markdownEditorJsx,
			markdowneditor.MarkdownEditor(),
		),

		r.Hr(nil),

		p.renderExample(
			exampleLatency,
			r.S("Latency Checker"),
			r.P(nil,
				r.S("By kind permission of "), r.A(&r.AProps{Href: "http://tjholowaychuk.com/"}, r.S("TJ Holowaychuk")),
				r.S(", a basic, component-based version of the beautiful APEX "), r.A(&r.AProps{Href: "https://latency.apex.sh/"}, r.S("Latency Tool")),
				r.S(" that uses randomly generated latency values."),
			),
			latencyJsx,
			r.A(&r.AProps{Href: "../latency", Target: "_blank"}, r.S("Launch in new tab")),
		),
	)

	return r.Div(&r.DivProps{ClassName: "container"},
		dc...,
	)
}

func (p ExamplesDef) renderExample(key exampleKey, title, msg r.Element, jsxSrc string, elem r.Element) r.Element {

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

func (p ExamplesDef) buildExampleNavTab(key exampleKey, t tab, title string) *r.LiElem {
	lip := &r.LiProps{Role: "presentation"}

	if v, _ := p.State().selectedTabs.Get(key); v == t {
		lip.ClassName = "active"
	}

	return r.Li(
		lip,
		r.A(
			&r.AProps{Href: "#", OnClick: tabChange{p, key, t}},
			r.S(title),
		),
	)

}

type tabChange struct {
	e   ExamplesDef
	key exampleKey
	t   tab
}

func (tc tabChange) OnClick(e *r.SyntheticMouseEvent) {
	p := tc.e
	key := tc.key
	t := tc.t

	cts := p.State().selectedTabs
	newSt := p.State()

	newSt.selectedTabs = cts.Set(key, t)
	p.SetState(newSt)

	e.PreventDefault()
}

func plainPanel(children ...r.Element) r.Element {
	return r.Div(&r.DivProps{ClassName: "panel panel-default panel-body"},
		children...,
	)
}
