// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

/*

Package react is a set of GopherJS bindings for Facebook's React, a Javascript
library for building user interfaces.

For more information see https://github.com/myitcv/react/wiki

*/
package react // import "myitcv.io/react"

import (
	"reflect"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"

	_ "myitcv.io/react/internal/bundle"
)

const (
	reactCompProps                     = "props"
	reactCompLastState                 = "_lastState"
	reactCompDisplayName               = "displayName"
	reactCompSetState                  = "setState"
	reactCompGetInitialState           = "getInitialState"
	reactCompShouldComponentUpdate     = "shouldComponentUpdate"
	reactCompComponentDidMount         = "componentDidMount"
	reactCompComponentWillReceiveProps = "componentWillReceiveProps"
	reactCompComponentWillMount        = "componentWillMount"
	reactCompComponentWillUnmount      = "componentWillUnmount"
	reactCompRender                    = "render"

	reactCreateElement = "createElement"
	reactCreateClass   = "createClass"
	reactDOMRender     = "render"

	nestedProps            = "_props"
	nestedState            = "_state"
	nestedComponentWrapper = "__ComponentWrapper"
)

var react = js.Global.Get("React")
var reactDOM = js.Global.Get("ReactDOM")
var object = js.Global.Get("Object")

// ComponentDef is embedded in a type definition to indicate the type is a component
type ComponentDef struct {
	elem *js.Object
	this *js.Object
}

var compMap = make(map[reflect.Type]*js.Object)

// S is the React representation of a string
type S string

func (s S) reactElement() {}

type Element interface {
	reactElement()
}

type generatesElement interface {
	element() *js.Object
}

type Component interface {
	ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool
	Render() Element

	setThis(this *js.Object)
	setElem(elem *js.Object)
}

type ComponentWithWillMount interface {
	Component
	ComponentWillMount()
}

type ComponentWithDidMount interface {
	Component
	ComponentDidMount()
}

type ComponentWithWillReceiveProps interface {
	Component
	ComponentWillReceivePropsIntf(i interface{})
}

type ComponentWithGetInitialState interface {
	Component
	GetInitialStateIntf() State
}

type ComponentWithWillUnmount interface {
	Component
	ComponentWillUnmount()
}

type State interface {
	IsState()
	EqualsIntf(v interface{}) bool
}

type Equals interface {
	EqualsIntf(v interface{}) bool
}

func (c *ComponentDef) reactElement() {}

func (c *ComponentDef) element() *js.Object {
	return c.elem
}

func (c *ComponentDef) Props() interface{} {
	if c.this != nil {
		return c.this.Get(reactCompProps).Get(nestedProps).Interface()
	}

	return c.elem.Get(reactCompProps).Get(nestedProps).Interface()
}

func (c *ComponentDef) SetState(i State) {
	cur := c.State()

	if i.EqualsIntf(cur) {
		return
	}

	res := object.New()
	res.Set(nestedState, js.MakeWrapper(i))
	c.this.Set(reactCompLastState, res)
	c.this.Call(reactCompSetState, res)
}

func (c *ComponentDef) setThis(this *js.Object) {
	c.this = this
}

func (c *ComponentDef) setElem(elem *js.Object) {
	c.elem = elem
}

func (c *ComponentDef) State() interface{} {
	return c.this.Get(reactCompLastState).Get(nestedState).Interface()
}

func BlessElement(cmp Component, newprops interface{}, children ...Element) {
	typ := reflect.TypeOf(cmp)

	comp, ok := compMap[typ]
	if !ok {
		comp = buildReactComponent(typ)
		compMap[typ] = comp
	}

	propsWrap := object.New()
	if newprops != nil {
		propsWrap.Set(nestedProps, js.MakeWrapper(newprops))
	}
	propsWrap.Set(nestedComponentWrapper, js.MakeWrapper(cmp))

	args := []interface{}{comp, propsWrap}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	elem := react.Call(reactCreateElement, args...)

	cmp.setElem(elem)
}

func buildReactComponent(typ reflect.Type) *js.Object {
	compDef := object.New()
	compDef.Set(reactCompDisplayName, typ.String())

	compDef.Set(reactCompGetInitialState, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		if cmp, ok := cw.Interface().(ComponentWithGetInitialState); ok {
			x := cmp.GetInitialStateIntf()
			if x == nil {
				return nil
			}
			res := object.New()
			res.Set(nestedState, js.MakeWrapper(x))
			this.Set(reactCompLastState, res)
			return res
		}

		return nil
	}))

	compDef.Set(reactCompShouldComponentUpdate, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		var nextProps interface{}
		var prevState interface{}
		var nextState interface{}

		if arguments[0] != nil {
			if i := arguments[0].Get(nestedProps); i != nil {
				nextProps = i.Interface()
			}
		}

		if arguments[1] != nil {
			if i := arguments[1].Get(nestedState); i != nil {
				nextState = i.Interface()
			}
		}

		if this != nil {
			if s := this.Get("state"); s != nil {
				if v := s.Get(nestedState); v != nil {
					prevState = v.Interface()
				}
			}
		}

		return cmp.ShouldComponentUpdateIntf(nextProps, prevState, nextState)
	}))

	compDef.Set(reactCompComponentDidMount, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		if cmp, ok := cw.Interface().(ComponentWithDidMount); ok {
			cmp.ComponentDidMount()
		}

		return nil
	}))

	compDef.Set(reactCompComponentWillReceiveProps, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		if cmp, ok := cw.Interface().(ComponentWithWillReceiveProps); ok {
			ourProps := arguments[0].Get(nestedProps).Interface()
			cmp.ComponentWillReceivePropsIntf(ourProps)
		}

		return nil
	}))

	compDef.Set(reactCompComponentWillUnmount, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		if cmp, ok := cw.Interface().(ComponentWithWillUnmount); ok {
			cmp.ComponentWillUnmount()
		}

		return nil
	}))

	compDef.Set(reactCompComponentWillMount, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		// TODO we can make this more efficient by not doing the type check
		// within the function body; it is known at the time of setting
		// "componentWillMount" on the compDef
		if cmp, ok := cmp.(ComponentWithWillMount); ok {
			cmp.ComponentWillMount()
		}

		return nil
	}))

	compDef.Set(reactCompRender, js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		props := this.Get(reactCompProps)
		cw := props.Get(nestedComponentWrapper)
		cmp := cw.Interface().(Component)

		cmp.setThis(this)

		renderRes := cmp.Render()

		return elementToReactObj(renderRes)
	}))

	return react.Call(reactCreateClass, compDef)
}

func elementToReactObj(el Element) interface{} {
	if el, ok := el.(generatesElement); ok {
		return el.element()
	}

	return el
}

func Render(el Element, container dom.Element) {
	reactDOM.Call(reactDOMRender, elementToReactObj(el), container)
}
