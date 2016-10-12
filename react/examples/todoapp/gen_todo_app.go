package todoapp

import "github.com/myitcv/gopherjs/react"

// Generated code (or at least will be once I write the code generator)
// for TodoApp

func (p *TodoAppDef) SetState(s TodoAppState) {
	p.ComponentDef.SetState(s)
}

func (p *TodoAppDef) State() TodoAppState {
	return p.ComponentDef.State().(TodoAppState)
}

func (p *TodoAppDef) GetInitialStateIntf() react.State {
	return p.GetInitialState()
}

func (p TodoAppState) IsState() {}
