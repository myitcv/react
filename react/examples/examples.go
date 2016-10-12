package main

import . "github.com/myitcv/gopherjs/react"

type ExamplesDef struct {
	ComponentDef
}

type tab int

const (
	tabJsx tab = iota
	tabGo
)

func Examples() *ExamplesDef {
	res := &ExamplesDef{}

	BlessElement(res, nil)

	return res
}

type ExamplesState struct {
	goSource     []string
	selectedTabs []tab
}

// // TODO fix Timer interaction bug before uncommenting
// func (p *ExamplesDef) ComponentWillMount() {
// 	for i, e := range examples {
// 		go func(i int, url string) {
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				panic(err)
// 			}

// 			defer resp.Body.Close()

// 			byts, err := ioutil.ReadAll(resp.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			newSt := p.State()
// 			newSt.goSource = make([]string, len(examples))
// 			copy(newSt.goSource, p.State().goSource)
// 			newSt.goSource[i] = string(byts)

// 			p.SetState(newSt)

// 		}(i, e.goSource)
// 	}
// }

func (p *ExamplesDef) GetInitialState() ExamplesState {
	return ExamplesState{
		goSource:     make([]string, len(examples)),
		selectedTabs: make([]tab, len(examples)),
	}
}

func (p *ExamplesDef) Render() Element {
	toRender := []Element{
		H3(nil, S("Reference")),
		P(nil,
			S("For the source code, raising issues, questions etc, please see "),
			A(
				AProps(func(ap *APropsDef) {
					ap.Href = "https://github.com/myitcv/gopherjs/tree/master/react/examples"
					ap.Target = "_blank"
				}),
				S("the Github repo"),
			),
		),
	}

	for i := range examples {
		if i > 0 {
			toRender = append(toRender, HR(nil))
		}

		toRender = append(toRender, p.renderExample(i))
	}

	return Div(
		DivProps(func(dp *DivPropsDef) {
			dp.ClassName = "container"
		}),

		toRender...,
	)
}

func (p *ExamplesDef) renderExample(i int) Element {
	e := examples[i]

	var code Element
	switch p.State().selectedTabs[i] {
	case tabGo:
		code = Code(
			CodeProps(func(cp *CodePropsDef) {
				cp.ClassName = "go"
			}),
			S("// Does not work for now"),
			// S(p.State().goSource[i]),
		)
	case tabJsx:
		code = Code(
			CodeProps(func(cp *CodePropsDef) {
				cp.ClassName = "nohighlight"
			}),
			S(e.jsxSource),
		)
	}

	return Div(nil,
		H3(nil, S(e.title)),
		Div(
			DivProps(func(dp *DivPropsDef) {
				dp.ClassName = "row"
			}),
			Div(
				DivProps(func(dp *DivPropsDef) {
					dp.ClassName = "col-md-8"
				}),
				Div(
					DivProps(func(dp *DivPropsDef) {
						dp.ClassName = "panel panel-default with-nav-tabs"
					}),
					Div(
						DivProps(func(dp *DivPropsDef) {
							dp.ClassName = "panel-heading"
						}),
						Ul(
							UlProps(func(ulp *UlPropsDef) {
								ulp.ClassName = "nav nav-tabs"
							}),

							p.buildExampleNavTab(i, tabJsx, "JSX"),
							// p.buildExampleNavTab(i, tabGo, "GopherJS (to follow)"),
						),
					),
					Div(
						DivProps(func(dp *DivPropsDef) {
							dp.ClassName = "panel-body"
						}),
						Pre(nil, code),
					),
				),
			),
			Div(
				DivProps(func(dp *DivPropsDef) {
					dp.ClassName = "col-md-4"
				}),
				plainPanel(
					e.elem(),
				),
			),
		),
	)
}

func (p *ExamplesDef) buildExampleNavTab(i int, t tab, title string) *LiDef {
	return Li(
		LiProps(func(lip *LiPropsDef) {
			if p.State().selectedTabs[i] == t {
				lip.ClassName = "active"
			}
			lip.Role = "presentation"
		}),
		A(
			AProps(func(ap *APropsDef) {
				ap.Href = "#"

				// TODO bug when clicking causes timer to blow up
				// ap.OnClick = p.handleTabChange(i, t)
			}),
			S(title),
		),
	)

}

func (p *ExamplesDef) handleTabChange(i int, t tab) func(*SyntheticMouseEvent) {
	return func(e *SyntheticMouseEvent) {
		newSt := p.State()

		// TODO this is a hack for now... should have to copy slice
		// when using PureRender
		newSt.selectedTabs[i] = t

		p.SetState(newSt)
	}
}

func plainPanel(children ...Element) Element {
	return Div(
		DivProps(func(dp *DivPropsDef) {
			dp.ClassName = "panel panel-default panel-body"
		}),
		children...,
	)
}
