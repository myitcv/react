package timer

import "github.com/myitcv/gopherjs/react"

// Generated code (or at least will be once I write the code generator)
// for Timer

func (p *TimerDef) SetState(s TimerState) {
	p.ComponentDef.SetState(s)
}

func (p *TimerDef) State() TimerState {
	return p.ComponentDef.State().(TimerState)
}

func (p *TimerDef) GetInitialStateIntf() react.State {
	return p.GetInitialState()
}

func (p TimerState) IsState() {}
