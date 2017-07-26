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

func App() *AppElem {
	return &AppElem{Element: r.CreateElement(buildApp, nil)}
}

func (a AppDef) Render() r.Element {
	return Latency()
}
