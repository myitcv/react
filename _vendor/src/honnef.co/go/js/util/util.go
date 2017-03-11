// Package util provides some helpers for working with GopherJS.
package util // import "honnef.co/go/js/util"

import "github.com/gopherjs/gopherjs/js"

func Float64Slice(o *js.Object) []float64 {
	if o == nil {
		return nil
	}
	d := o.Interface().([]interface{})
	ret := make([]float64, len(d))
	for i, e := range d {
		ret[i] = e.(float64)
	}
	return ret
}

func IntSlice(o *js.Object) []int {
	if o == nil {
		return nil
	}
	d := o.Interface().([]interface{})
	ret := make([]int, len(d))
	for i, e := range d {
		ret[i] = int(e.(float64))
	}
	return ret
}

func StringSlice(o *js.Object) []string {
	if o == nil {
		return nil
	}
	d := o.Interface().([]interface{})
	ret := make([]string, len(d))
	for i, e := range d {
		ret[i] = e.(string)
	}
	return ret
}

type Err struct {
	*js.Object
	Message string `js:"message"`
	Name    string `js:"name"`
	File    string `js:"fileName"`   // Mozilla extension
	Line    int    `js:"lineNumber"` // Mozilla extension
	Stack   string `js:"stack"`      // Chrome/Microsoft extension
}

func (err Err) Error() string {
	return err.Message
}

func Error(o *js.Object) error {
	if o == nil {
		return nil
	}
	return Err{Object: o}
}

type EventTarget struct {
	*js.Object
}

func (t EventTarget) AddEventListener(typ string, useCapture bool, listener func(*js.Object)) {
	t.Call("addEventListener", typ, listener, useCapture)
}

func (t EventTarget) RemoveEventListener(typ string, useCapture bool, listener func(*js.Object)) {
	t.Call("removeEventListener", typ, listener, useCapture)
}
