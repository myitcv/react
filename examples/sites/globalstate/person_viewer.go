package main

import (
	"fmt"

	r "myitcv.io/react"
	"myitcv.io/react/examples/sites/globalstate/model"
	"myitcv.io/react/examples/sites/globalstate/state"
)

type PersonViewerDef struct {
	r.ComponentDef
}

type PersonViewerState struct {
	p *model.Person

	curPersSub *state.Sub
}

func PersonViewer() *PersonViewerElem {
	return buildPersonViewerElem()
}

func (p PersonViewerDef) ComponentWillMount() {
	curPersSub := state.State.CurrentPerson().Subscribe(p.currPersonUpdated)

	p.SetState(PersonViewerState{
		p:          state.State.CurrentPerson().Get(),
		curPersSub: curPersSub,
	})
}

func (p PersonViewerDef) ComponentWillUnmount() {
	p.State().curPersSub.Clear()
}

func (p PersonViewerDef) Render() r.Element {
	st := p.State()

	if st.p != nil {
		return r.P(nil, r.S(fmt.Sprintf("You have selected %v, age %v", st.p.Name(), st.p.Age())))
	}

	return r.P(nil, r.S("(no person selected)"))
}

func (p PersonViewerDef) currPersonUpdated() {
	st := p.State()
	st.p = state.State.CurrentPerson().Get()
	p.SetState(st)
}
