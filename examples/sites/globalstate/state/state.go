package state

import "myitcv.io/react/examples/sites/globalstate/model"

//go:generate stateGen

var State = NewRoot()

var root _Node_App

type _Node_App struct {
	CurrentPerson *model.Person
	Root          *_Node_Data
}

type _Node_Data struct {
	People *model.People
}
