package react

import "github.com/gopherjs/gopherjs/js"

// ADef is the React component definition corresponding to the HTML <a> element
type ADef struct {
	underlying *js.Object
}

// APropsDef defines the properties for the <a> element
type APropsDef struct {
	*BasicHTMLElement

	Target string `js:"target"`
	Href   string `js:"href"`
}

// AProps creates a new instance of <a> properties, mutating the props
// by the provided initiaaser
func AProps(f func(p *APropsDef)) *APropsDef {
	res := &APropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *ADef) reactElement() {}

// A creates a new instance of a <a> element with the provided props and
// children
func A(props *APropsDef, children ...Element) *ADef {
	args := []interface{}{"a", props}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &ADef{underlying: underlying}
}
