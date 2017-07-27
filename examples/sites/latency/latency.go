package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"honnef.co/go/js/dom"
	r "myitcv.io/react"
	"myitcv.io/react/jsx"
)

//go:generate reactGen
//go:generate immutableGen

type location string

var (
	locations = [...]location{
		"Oregon",
		"California",
		"Ohio",
		"Virginia",
		"Ireland",
		"Frankfurt",
		"London",
		"Mumbai",
		"Singapore",
		"Seoul",
		"Tokyo",
		"Sydney",
	}
)

type _Imm_latencies map[location]latency

type latency struct {
	dns      time.Duration
	tcp      time.Duration
	tls      time.Duration
	wait     time.Duration
	download time.Duration
	total    time.Duration
}

type LatencyDef struct {
	r.ComponentDef
}

type LatencyState struct {
	reqId  int
	output bool

	url    string
	altUrl string

	*latencies
}

func Latency() *LatencyElem {
	return buildLatencyElem()
}

func (l LatencyDef) Render() r.Element {
	var c r.Element

	if l.State().output {
		c = l.renderOutput()
	} else {
		c = l.renderInput()
	}

	return r.Div(&r.DivProps{ClassName: "App"},
		r.Div(&r.DivProps{ClassName: "Content center full column"},
			jsx.HTMLElem(`
			<div className="Title margin center">
				<span className="text">Latency</span>
				<span className="subtext">Global latency testing tool</span>
			</div>
			`),
			c,
			jsx.HTMLElem(`
			<div className="Title margin center">
				<br/>
				<span className="subtext" style="font-size:smaller; font-style: italic">(randomly generated results)</span>
				<span className="subtext" style="font-size:smaller; font-style: italic">
					Real, original version <a href="https://latency.apex.sh/" target="_blank">https://latency.apex.sh/</a>
				</span>
			</div>
			`),
		),
	)
}

func (l LatencyDef) renderInput() r.Element {
	return r.Form(&r.FormProps{ClassName: "LatencyForm"},
		r.Div(&r.DivProps{ClassName: "group"},
			r.Input(&r.InputProps{
				Type:        "text",
				ID:          "url",
				Placeholder: "url to test (can be anything)",
				Value:       l.State().url,
				OnChange:    urlChange{l},
			}),
			r.Input(&r.InputProps{
				Type:        "text",
				ID:          "altUrl",
				Placeholder: "comparison url (not used)",
				Value:       l.State().altUrl,
				OnChange:    altUrlChange{l},
			}),
		),
		r.Button(
			&r.ButtonProps{
				ClassName: "Button small",
				OnClick:   check{l},
			},
			r.S("Start"),
		),
	)
}

const (
	// gross hack for now
	resultWidth = 500.0
)

func (l *LatencyDef) renderOutput() r.Element {
	var regions []r.Element

	ls := l.State().latencies

	maxTot := time.Duration(0)

	for _, lat := range ls.Range() {
		if lat.total > maxTot {
			maxTot = lat.total
		}
	}

	awfulTime := maxTot / 3
	okTime := maxTot * 2 / 3

	for _, v := range locations {
		regClass := "Region"

		timings := []r.Element{
			r.Span(&r.SpanProps{ClassName: "total"}, r.S("0ms")),
		}

		res, ok := ls.Get(v)
		if ok {
			if res.total < awfulTime {
				regClass += " with-timings speed-awful"
			} else if res.total < okTime {
				regClass += " with-timings speed-ok"
			} else {
				regClass += " with-timings speed-fast"
			}

			genTiming := func(f time.Duration, n, l string) *r.SpanElem {
				w := fmt.Sprintf("%.3fpx", float64(f)/float64(maxTot)*resultWidth)
				rs := fmt.Sprintf("%v (%v)", l, f)

				return r.Span(
					&r.SpanProps{
						ClassName: "timing " + n,
						Style:     &r.CSS{Width: w},
					},
					r.S(rs),
				)
			}

			timings = []r.Element{
				genTiming(res.dns, "dns", "DNS"),
				genTiming(res.tcp, "tcp", "TCP"),
				genTiming(res.tls, "tls", "TLS"),
				genTiming(res.wait, "wait", "Wait"),
				genTiming(res.download, "download", "Download"),
				r.Span(&r.SpanProps{ClassName: "total"}, r.S(fmt.Sprintf("%v", res.total))),
			}
		}

		rd := r.Div(&r.DivProps{ClassName: regClass},
			r.Span(&r.SpanProps{ClassName: "name"}, r.S(v)),
			r.Div(&r.DivProps{ClassName: "Results"},
				r.Div(&r.DivProps{ClassName: "Timings"}, timings...),
			),
		)

		regions = append(regions, rd)
	}

	return r.Div(&r.DivProps{ClassName: "Regions"},
		regions...,
	)
}

func (l *LatencyDef) reset(e *r.SyntheticMouseEvent) {
	s := l.State()
	s.output = false
	s.reqId++
	l.SetState(s)

	e.PreventDefault()
}

type urlChange struct{ l LatencyDef }
type altUrlChange struct{ l LatencyDef }
type check struct{ l LatencyDef }

func (u urlChange) OnChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)
	s := u.l.State()
	s.url = target.Value
	u.l.SetState(s)
}

func (a altUrlChange) OnChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)
	s := a.l.State()
	s.altUrl = target.Value
	a.l.SetState(s)
}

// this could clearly be replace by something that actually checks
// the realy latencies instead of randomly generating them
func (c check) OnClick(e *r.SyntheticMouseEvent) {
	l := c.l

	reqId := l.State().reqId

	for _, v := range locations {
		loc := v
		to := rand.Int63n(3000) * int64(time.Millisecond)

		go func() {
			<-time.After(time.Duration(to))
			s := l.State()

			if s.reqId == reqId {
				lat := latency{}

				ints := make([]int64, 4)

				for i := range ints {
					ints[i] = rand.Int63n(to)
				}

				ints = append([]int64{0}, ints...)
				ints = append(ints, to)

				sort.Slice(ints, func(i, j int) bool {
					return ints[i] < ints[j]
				})

				vs := make([]time.Duration, len(ints)-1)

				for i := range vs {
					vs[i] = time.Duration(ints[i+1] - ints[i])
				}

				lat.dns = vs[0]
				lat.tcp = vs[1]
				lat.tls = vs[2]
				lat.wait = vs[3]
				lat.download = vs[4]
				lat.total = time.Duration(to)

				s.latencies = s.latencies.Set(loc, lat)

				l.SetState(s)
			}

		}()
	}

	s := l.State()
	s.output = true
	s.latencies = newLatencies()
	l.SetState(s)

	e.PreventDefault()
}
