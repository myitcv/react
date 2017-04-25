package main

import (
	"math/rand"
	"time"

	r "myitcv.io/react"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
