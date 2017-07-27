// foo_bar.go

package main

import (
	"fmt"

	// it's normal to name the import of myitcv.io/react with something like
	// "r" because it's used repeatedly. Dot imports are strongly discouraged
	// per the core Go team's advice on this subject.
	//
	r "myitcv.io/react"
)

//go:generate reactGen

// FooBarDef is the definition of the FooBar component. All components are
// declared with a *Def suffix and an embedded myitcv.io/react.ComponentDef
// field
//
type FooBarDef struct {
	r.ComponentDef
}

// FooBarProps is the props type for the FooBar component. All props types are
// declared as a struct type with a *Props suffix
//
type FooBarProps struct {
	Name string
}

// FooBarState is the state type for the FooBar component. All state types are
// declared as a struct type with a *State suffix
//
type FooBarState struct {
	Age int
}

// FooBar is the constructor for a FooBar component. Given that this component
// can take props (can, not must), we add a parameter of type FooBarProps
//
func FooBar(p FooBarProps) *FooBarElem {
	// every component constructor must call this function
	return buildFooBarElem(p)
}

// Render is a required method on all React components. Notice that the method
// is declared on the type FooBarDef.
//
func (f FooBarDef) Render() r.Element {

	// all React components must render under a single root. This is typically achieved
	// by rendering everything within a <div> elememt
	//
	return r.Div(nil,
		r.P(nil,
			r.S(fmt.Sprintf("My name is %v. My age is %v", f.Props().Name, f.State().Age)),
		),
		r.Button(
			&r.ButtonProps{
				OnClick: ageClick{f},
			},
			r.S("Bump age"),
		),
	)
}

// ageClick implements the react.OnClick interface to handle when the "Bump age" button
// is clicked
//
type ageClick struct{ FooBarDef }

// OnClick is the ageClick implementation of the react.OnClick interface
//
func (a ageClick) OnClick(e *r.SyntheticMouseEvent) {
	f := a.FooBarDef

	s := f.State()
	s.Age++
	f.SetState(s)
}
