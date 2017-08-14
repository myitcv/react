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
func HelloMessage(p HelloMessageProps) *HelloMessageElem {
	return buildHelloMessageElem(p)
}

// Render renders the HelloMessage component
func (h HelloMessageDef) Render() react.Element {
	return react.Div(nil,
		react.S("Hello "+h.Props().Name),
	)
}
