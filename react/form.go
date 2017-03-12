package react

import "github.com/gopherjs/gopherjs/js"

// FormDef is the React component definition corresponding to the HTML <form> element
type FormDef struct {
	underlying *js.Object
}

// FormPropsDef defines the properties for the <form> element
type FormPropsDef struct {
	*BasicHTMLElement
}

// FormProps creates a new instance of <form> properties, mutating the props
// by the provided initiaformser
func FormProps(f func(p *FormPropsDef)) *FormPropsDef {
	res := &FormPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *FormDef) reactElement() {}

// Form creates a new instance of a <form> element with the provided props and
// children
func Form(props *FormPropsDef, children ...Element) *FormDef {
	args := []interface{}{"form", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &FormDef{underlying: underlying}
}
