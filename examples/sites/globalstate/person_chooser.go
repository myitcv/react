package main

import (
	"myitcv.io/react"
	"myitcv.io/react/components/imm"
	"myitcv.io/react/examples/sites/globalstate/model"
	"myitcv.io/react/examples/sites/globalstate/state"
	"myitcv.io/sorter"
)

//go:generate sortGen

type PersonState interface {
	Get() *model.Person
	Set(p *model.Person)
	Subscribe(cb func()) *state.Sub
}

type PersonChooserDef struct {
	react.ComponentDef
}

type PersonChooserProps struct {
	PersonState
}

type PersonChooserState struct {
	currPerson    *model.Person
	currPersonSub *state.Sub
}

func PersonChooser(props PersonChooserProps) *PersonChooserElem {
	return buildPersonChooserElem(props)
}

func (p PersonChooserDef) ComponentWillMount() {
	sub := p.Props().PersonState.Subscribe(p.personStateChanged)
	st := p.State()
	st.currPersonSub = sub
	st.currPerson = p.Props().PersonState.Get()
	p.SetState(st)
}

func (p PersonChooserDef) ComponentWillUnmount() {
	p.State().currPersonSub.Clear()
}

func (p PersonChooserDef) Render() react.Element {

	ppl := sortPeopleKeysByName(state.State.Root().People().Get())

	ps := []imm.Label{personLabel{nil}}

	for _, v := range ppl.Range() {
		ps = append(ps, personLabel{v})
	}

	return imm.Select(
		imm.SelectProps{
			Entry:    personLabel{p.State().currPerson},
			Entries:  imm.NewLabelEntries(ps...),
			OnSelect: personSelected{p},
		},
	)
}

func (p PersonChooserDef) personStateChanged() {
	s := p.State()
	s.currPerson = p.Props().PersonState.Get()
	p.SetState(s)
}

func orderPeopleKeysByName(ppl *model.People, i, j int) sorter.Ordered {
	lhs := ppl.Get(i)
	rhs := ppl.Get(j)

	return lhs.Name() < rhs.Name()
}

type personLabel struct{ *model.Person }

func (p personLabel) Label() string {
	if p.Person == nil {
		return ""
	}

	return p.Person.Name()
}

type personSelected struct{ PersonChooserDef }

func (p personSelected) OnSelect(l imm.Label) {
	pl := l.(personLabel)
	s := p.PersonChooserDef.State()
	s.currPerson = pl.Person
	p.PersonChooserDef.SetState(s)

	p.PersonChooserDef.Props().Set(pl.Person)
}
