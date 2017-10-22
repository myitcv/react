package hellomessage // import "myitcv.io/react/examples/hellomessage"

import (
	"myitcv.io/react"
)

//go:generate reactGen

// HelloMessageDef is the definition of the HelloMessage component
type HelloMessageDef struct {
	react.ComponentDef
}

// HelloMessageProps is the props type for the HelloMessage component
type HelloMessageProps struct {
	Name string
}

// HelloMessage creates instances of the HelloMessage component
func HelloMessage(p HelloMessageProps, children ...react.Element) *HelloMessageElem {
	return buildHelloMessageElem(p, children...)
}

// Render renders the HelloMessage component
func (h HelloMessageDef) Render() *react.DivElem {
	kids := []react.Element{react.S("Hello " + h.Props().Name)}

	for _, v := range h.Children() {
		kids = append(kids, v)
	}

	return react.Div(nil,
		kids...,
	)
}

func (h HelloMessageDef) RendersDiv(*react.DivElem) {}
