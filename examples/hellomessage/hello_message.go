package hellomessage // import "myitcv.io/react/examples/hellomessage"

import (
	r "myitcv.io/react"
)

//go:generate reactGen

// HelloMessageDef is the definition of the HelloMessage component
type HelloMessageDef struct {
	r.ComponentDef
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
func (h HelloMessageDef) Render() r.Element {
	return r.Div(nil,
		r.S("Hello "+h.Props().Name),
	)
}
