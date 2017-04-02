package examples

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples/immtodoapp"
	"honnef.co/go/js/xhr"
)

// ImmExamplesDef is the definition of the ImmExamples component
type ImmExamplesDef struct {
	r.ComponentDef
}

// ImmExamples creates instances of the ImmExamples component
func ImmExamples() *ImmExamplesDef {
	res := new(ImmExamplesDef)
	r.BlessElement(res, nil)
	return res
}

// ImmExamplesState is the state type for the ImmExamples component
type ImmExamplesState struct {
	examples     *exampleSource
	selectedTabs *tabS
}

// ComponentWillMount is a React lifecycle method for the ImmExamples component
func (p *ImmExamplesDef) ComponentWillMount() {
	if !fetchStarted {
		for i, e := range sources.Range() {
			go func(i exampleKey, e *source) {
				req := xhr.NewRequest("GET", "https://raw.githubusercontent.com/myitcv/gopherjs/master/react/examples/"+e.file())
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
func (p *ImmExamplesDef) GetInitialState() ImmExamplesState {
	return ImmExamplesState{
		examples:     sources,
		selectedTabs: newTabS(),
	}
}

// Render renders the ImmExamples component
func (p *ImmExamplesDef) Render() r.Element {
	toRender := []r.Element{
		r.H3(nil, r.S("Reference")),
		r.P(nil, r.S("This entire page is a React application. An outer "), r.Code(nil, r.S("ImmExamples")), r.S(" component contains a number of inner components.")),
		r.P(nil,
			r.S("For the source code, raising issues, questions etc, please see "),
			r.A(
				&r.AProps{
					Href:   "https://github.com/myitcv/gopherjs/tree/master/react/examples",
					Target: "_blank",
				},
				r.S("the Github repo"),
			),
			r.S("."),
		),
		r.P(nil,
			r.S("Note the examples below show the GopherJS source code from "), r.Code(nil, r.S("master")),
		),

		p.renderExample(
			exampleImmTodo,
			r.Span(nil, r.S("An Application using "), r.Code(nil, r.S("github.com/myitcv/immutable"))),
			r.P(nil, r.S("The immtodoapp.TodoApp component is a reimplementation of todoapp.TodoApp using immutable data structures.")),
			"n/a",
			immtodoapp.TodoApp(),
		),
	}

	return r.Div(&r.DivProps{ClassName: "container"},
		toRender...,
	)
}

func (p *ImmExamplesDef) renderExample(key exampleKey, title, msg r.Element, jsxSrc string, elem r.Element) r.Element {

	var goSrc string
	src, _ := p.State().examples.Get(key)
	if src != nil {
		goSrc = src.src()
	}

	var code r.Element
	switch v, _ := p.State().selectedTabs.Get(key); v {
	case tabGo:
		code = r.Pre(nil, r.S(goSrc))
	case tabJsx:
		code = r.Pre(nil, r.S(jsxSrc))
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
						r.Pre(nil, code),
					),
				),
			),
			r.Div(&r.DivProps{ClassName: "col-md-4"},
				plainPanel(elem),
			),
		),
	)
}

func (p *ImmExamplesDef) buildExampleNavTab(key exampleKey, t tab, title string) *r.LiDef {
	lip := &r.LiProps{Role: "presentation"}

	if v, _ := p.State().selectedTabs.Get(key); v == t {
		lip.ClassName = "active"
	}

	return r.Li(
		lip,
		r.A(
			&r.AProps{Href: "#", OnClick: p.handleTabChange(key, t)},
			r.S(title),
		),
	)

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
