package hellomessage

import (
	. "github.com/myitcv/gopherjs/react"
)

type HelloMessageDef struct {
	ComponentDef
}

type HelloMessageProps struct {
	Name string
}

func HelloMessage(p HelloMessageProps) *HelloMessageDef {
	res := &HelloMessageDef{}

	BlessElement(res, p)

	return res
}

func (r *HelloMessageDef) Render() Element {
	return Div(nil,
		S("Hello "+r.Props().Name),
	)
}
