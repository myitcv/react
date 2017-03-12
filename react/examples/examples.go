package main

import (
	"honnef.co/go/js/xhr"

	r "github.com/myitcv/gopherjs/react"
)

type ExamplesDef struct {
	r.ComponentDef
}

type tab int

const (
	tabGo tab = iota
	tabJsx
)

func Examples() *ExamplesDef {
	res := &ExamplesDef{}

	r.BlessElement(res, nil)

	return res
}

type ExamplesState struct {
	goSource     []string
	selectedTabs []tab
}

// // TODO fix Timer interaction bug before uncommenting
func (p *ExamplesDef) ComponentDidMount() {
	for i, e := range examples {
		go func(i int, url string) {
			req := xhr.NewRequest("GET", "https://raw.githubusercontent.com/myitcv/gopherjs/master/react/examples/"+url)
			err := req.Send(nil)
			if err != nil {
				panic(err)
			}

			newSt := p.State()
			newSt.goSource = make([]string, len(examples))
			copy(newSt.goSource, p.State().goSource)
			newSt.goSource[i] = req.ResponseText

			p.SetState(newSt)

		}(i, e.goSource)
	}
}

func (p *ExamplesDef) GetInitialState() ExamplesState {
	return ExamplesState{
		goSource:     make([]string, len(examples)),
		selectedTabs: make([]tab, len(examples)),
	}
}

func (p *ExamplesDef) Render() r.Element {
	toRender := []r.Element{
		r.H3(nil, r.S("Reference")),
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

	for i := range examples {
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
	e := examples[i]

	var code r.Element
	switch p.State().selectedTabs[i] {
	case tabGo:
		code = r.Pre(nil, r.S(p.State().goSource[i]))
	case tabJsx:
		code = r.Pre(nil, r.S(e.jsxSource))
	}

	return r.Div(nil,
		r.H3(nil, r.S(e.title)),
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
					e.elem(),
				),
			),
		),
	)
}

func (p *ExamplesDef) buildExampleNavTab(i int, t tab, title string) *r.LiDef {
	return r.Li(
		r.LiProps(func(lip *r.LiPropsDef) {
			if p.State().selectedTabs[i] == t {
				lip.ClassName = "active"
			}
			lip.Role = "presentation"
		}),
		r.A(
			r.AProps(func(ap *r.APropsDef) {
				ap.Href = "#"

				// TODO fix Timer interaction bug before uncommenting
				ap.OnClick = p.handleTabChange(i, t)
			}),
			r.S(title),
		),
	)

}

func (p *ExamplesDef) handleTabChange(i int, t tab) func(*r.SyntheticMouseEvent) {
	return func(e *r.SyntheticMouseEvent) {
		newSt := p.State()

		// TODO this is a hack for now... should have to copy slice
		// when using PureRender
		newSt.selectedTabs[i] = t

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
