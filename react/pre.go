package react

import "github.com/gopherjs/gopherjs/js"

// PreDef is the React component definition corresponding to the HTML <pre> element
type PreDef struct {
	underlying *js.Object
}

// PrePropsDef defines the properties for the <pre> element
type PrePropsDef struct {
	*BasicHTMLElement
}

// PreProps creates a new instance of <pre> properties, mutating the props
// by the provided initiapreser
func PreProps(f func(p *PrePropsDef)) *PrePropsDef {
	res := &PrePropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *PreDef) reactElement() {}

// Pre creates a new instance of a <pre> element with the provided props and
// children
func Pre(props *PrePropsDef, children ...Element) *PreDef {
	args := []interface{}{"pre", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &PreDef{underlying: underlying}
}
