package react

import "github.com/gopherjs/gopherjs/js"

// PDef is the React component definition corresponding to the HTML <p> element
type PDef struct {
	underlying *js.Object
}

// PPropsDef defines the properties for the <p> element
type PPropsDef struct {
	*BasicHTMLElement
}

// PProps creates a new instance of <p> properties, mutating the props
// by the provided initiapser
func PProps(f func(p *PPropsDef)) *PPropsDef {
	res := &PPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *PDef) reactElement() {}

// P creates a new instance of a <p> element with the provided props and
// children
func P(props *PPropsDef, children ...Element) *PDef {
	args := []interface{}{"p", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &PDef{underlying: underlying}
}
