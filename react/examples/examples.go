package examples

import (
	r "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/xhr"
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
func Examples() *ExamplesDef {
	res := new(ExamplesDef)
	r.BlessElement(res, nil)
	return res
}

type (
	_Imm_exampleS []*example
	_Imm_tabS     []tab
)

// ExamplesState is the state type for the Examples component
type ExamplesState struct {
	examples     *exampleS
	selectedTabs *tabS
}

// ComponentWillMount is a React lifecycle method for the Examples component
func (p *ExamplesDef) ComponentWillMount() {
	if !fetchStarted {
		for i, e := range examples.Range() {
			go func(i int, e *example) {
				req := xhr.NewRequest("GET", "https://raw.githubusercontent.com/myitcv/gopherjs/master/react/examples/"+e.goSourceFile())
				err := req.Send(nil)
				if err != nil {
					panic(err)
				}

				examples = examples.Set(i, e.setGoSourceStr(req.ResponseText))

				newSt := p.State()
				newSt.examples = examples
				p.SetState(newSt)
			}(i, e)
		}

		fetchStarted = true
	}
}

// GetInitialState returns in the initial state for the Examples component
func (p *ExamplesDef) GetInitialState() ExamplesState {
	return ExamplesState{
		examples:     examples,
		selectedTabs: newTabSLen(examples.Len()),
	}
}

// Render renders the Examples component
func (p *ExamplesDef) Render() r.Element {
	toRender := []r.Element{
		r.H3(nil, r.S("Reference")),
		r.P(nil, r.S("This entire page is a React application. An outer "), r.Code(nil, r.S("Examples")), r.S(" component contains a number of inner components.")),
		r.P(nil,
			r.S("For the source code, raising issues, questions etc, please see "),
			r.A(
				r.AProps(func(ap *r.APropsDef) {
					ap.Href = "https://github.com/myitcv/gopherjs/tree/master/react/examples"
					ap.Target = "_blank"
				}),
				r.S("the Github repo"),
			),
			r.S("."),
		),
		r.P(nil,
			r.S("Note the examples below show the GopherJS source code from "), r.Code(nil, r.S("master")),
		),
	}

	for i := range p.State().examples.Range() {
		if i > 0 {
			toRender = append(toRender, r.HR(nil))
		}

		toRender = append(toRender, p.renderExample(i))
	}

	return r.Div(
		r.DivProps(func(dp *r.DivPropsDef) {
			dp.ClassName = "container"
		}),

		toRender...,
	)
}

func (p *ExamplesDef) renderExample(i int) r.Element {
	e := p.State().examples.Get(i)

	var code r.Element
	switch p.State().selectedTabs.Get(i) {
	case tabGo:
		code = r.Pre(nil, r.S(p.State().examples.Get(i).goSourceStr()))
	case tabJsx:
		code = r.Pre(nil, r.S(e.jsxSourceStr()))
	}

	return r.Div(nil,
		r.H3(nil, r.S(e.title())),
		r.P(nil, r.S(e.message())),
		r.Div(
			r.DivProps(func(dp *r.DivPropsDef) {
				dp.ClassName = "row"
			}),
			r.Div(
				r.DivProps(func(dp *r.DivPropsDef) {
					dp.ClassName = "col-md-8"
				}),
				r.Div(
					r.DivProps(func(dp *r.DivPropsDef) {
						dp.ClassName = "panel panel-default with-nav-tabs"
					}),
					r.Div(
						r.DivProps(func(dp *r.DivPropsDef) {
							dp.ClassName = "panel-heading"
						}),
						r.Ul(
							r.UlProps(func(ulp *r.UlPropsDef) {
								ulp.ClassName = "nav nav-tabs"
							}),

							p.buildExampleNavTab(i, tabGo, "GopherJS"),
							p.buildExampleNavTab(i, tabJsx, "JSX"),
						),
					),
					r.Div(
						r.DivProps(func(dp *r.DivPropsDef) {
							dp.ClassName = "panel-body"
						}),
						r.Pre(nil, code),
					),
				),
			),
			r.Div(
				r.DivProps(func(dp *r.DivPropsDef) {
					dp.ClassName = "col-md-4"
				}),
				plainPanel(
					e.elem()(),
				),
			),
		),
	)
}

func (p *ExamplesDef) buildExampleNavTab(i int, t tab, title string) *r.LiDef {
	return r.Li(
		r.LiProps(func(lip *r.LiPropsDef) {
			if p.State().selectedTabs.Get(i) == t {
				lip.ClassName = "active"
			}
			lip.Role = "presentation"
		}),
		r.A(
			r.AProps(func(ap *r.APropsDef) {
				ap.Href = "#"
				ap.OnClick = p.handleTabChange(i, t)
			}),
			r.S(title),
		),
	)

}

func (p *ExamplesDef) handleTabChange(i int, t tab) func(*r.SyntheticMouseEvent) {
	return func(e *r.SyntheticMouseEvent) {
		cts := p.State().selectedTabs
		newSt := p.State()

		newSt.selectedTabs = cts.Set(i, t)
		p.SetState(newSt)

		e.PreventDefault()
	}
}

func plainPanel(children ...r.Element) r.Element {
	return r.Div(
		r.DivProps(func(dp *r.DivPropsDef) {
			dp.ClassName = "panel panel-default panel-body"
		}),
		children...,
	)
}
