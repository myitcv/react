package react

import "github.com/gopherjs/gopherjs/js"

// CodeDef is the React component definition corresponding to the HTML <code> element
type CodeDef struct {
	underlying *js.Object
}

// CodePropsDef defines the properties for the <code> element
type CodePropsDef struct {
	*BasicHTMLElement
}

// CodeProps creates a new instance of <code> properties, mutating the props
// by the provided initiacodeser
func CodeProps(f func(p *CodePropsDef)) *CodePropsDef {
	res := &CodePropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *CodeDef) reactElement() {}

// Code creates a new instance of a <code> element with the provided props
func Code(props *CodePropsDef, children ...Element) *CodeDef {
	args := []interface{}{"code", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &CodeDef{underlying: underlying}
}
