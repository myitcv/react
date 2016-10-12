package main

import "github.com/myitcv/gopherjs/react"

// Generated code (or at least will be once I write the code generator)
// for Examples

func (p *ExamplesDef) SetState(s ExamplesState) {
	p.ComponentDef.SetState(s)
}

func (p *ExamplesDef) State() ExamplesState {
	return p.ComponentDef.State().(ExamplesState)
}

func (p *ExamplesDef) GetInitialStateIntf() react.State {
	return p.GetInitialState()
}

func (p ExamplesState) IsState() {}
