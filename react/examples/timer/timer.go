package timer

import (
	"fmt"
	"time"

	r "github.com/myitcv/gopherjs/react"
)

//go:generate reactGen

type TimerDef struct {
	r.ComponentDef

	ticker *time.Ticker
}

type TimerState struct {
	secondsElapsed float32
}

func Timer() *TimerDef {
	res := &TimerDef{}

	r.BlessElement(res, nil)

	return res
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

func (p *TimerDef) Render() r.Element {
	return r.Div(nil,
		r.Div(nil,
			r.S(fmt.Sprintf("Seconds elapsed %.0f", p.State().secondsElapsed)),
		),
	)
}
