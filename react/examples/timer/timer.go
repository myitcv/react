// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package timer

import (
	"fmt"
	"time"

	r "github.com/myitcv/gopherjs/react"
)

//go:generate reactGen

// TimerDef is the definition of the Timer component
type TimerDef struct {
	r.ComponentDef

	ticker *time.Ticker
}

// TimerState is the state type for the Timer component
type TimerState struct {
	secondsElapsed float32
}

// Timer creates instances of the Timer component
func Timer() *TimerDef {
	res := new(TimerDef)
	r.BlessElement(res, nil)
	return res
}

// ComponentWillMount is a React lifecycle method for the Timer component
func (t *TimerDef) ComponentWillMount() {
	t.ticker = time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-t.ticker.C:
				c := t.State()
				c.secondsElapsed++
				t.SetState(c)
			}
		}
	}()
}

// Render renders the Timer component
func (t *TimerDef) Render() r.Element {
	return r.Div(nil,
		r.Div(nil,
			r.S(fmt.Sprintf("Seconds elapsed %.0f", t.State().secondsElapsed)),
		),
	)
}
