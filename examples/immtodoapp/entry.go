package immtodoapp

import "myitcv.io/react"

type EntryDef struct {
	react.ComponentDef
}

type EntryProps struct {
	Text string
}

func Entry(s string) *EntryElem {
	return buildEntryElem(EntryProps{Text: s})
}

func (e EntryDef) Render() *react.LiElem {
	return react.Li(nil, react.S(e.Props().Text))
}

func (e EntryDef) RendersLi(*react.LiElem) {}
