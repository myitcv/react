package main

import (
	r "myitcv.io/react"
)

type AppDef struct {
	r.ComponentDef
}

func App() *AppDef {
	res := new(AppDef)
	r.BlessElement(res, nil)
	return res
}

func (a *AppDef) Render() r.Element {
	return Latency()
}
