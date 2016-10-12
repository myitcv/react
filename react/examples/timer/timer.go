package timer

import (
	"fmt"
	"time"

	. "github.com/myitcv/gopherjs/react"
)

type TimerDef struct {
	ComponentDef

	ticker *time.Ticker
}

type TimerState struct {
	secondsElapsed float32
}

func Timer() *TimerDef {
	res := &TimerDef{}

	BlessElement(res, nil)

	return res
}

func (p *TimerDef) GetInitialState() TimerState {
	return TimerState{
		secondsElapsed: 0,
	}
}

func (p *TimerDef) ComponentWillUnmount() {
	p.ticker.Stop()
}

func (p *TimerDef) ComponentWillMount() {
	p.ticker = time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-p.ticker.C:
				c := p.State()
				c.secondsElapsed++
				p.SetState(c)
			}
		}
	}()
}

func (p *TimerDef) Render() Element {
	return Div(nil,
		Div(nil,
			S(fmt.Sprintf("Seconds elapsed %.0f", p.State().secondsElapsed)),
		),
	)
}
