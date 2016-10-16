package hellomessage

import (
	r "github.com/myitcv/gopherjs/react"
)

type HelloMessageDef struct {
	r.ComponentDef
}

type HelloMessageProps struct {
	Name string
}

func HelloMessage(p HelloMessageProps) *HelloMessageDef {
	res := &HelloMessageDef{}

	r.BlessElement(res, p)

	return res
}

func (h *HelloMessageDef) Render() r.Element {
	return r.Div(nil,
		r.S("Hello "+h.Props().Name),
	)
}
