package react

import "github.com/gopherjs/gopherjs/js"

// InputDef is the React component definition corresponding to the HTML <input> element
type InputDef struct {
	underlying *js.Object
}

// InputPropsDef defines the properties for the <input> element
type InputPropsDef struct {
	*BasicHTMLElement

	Placeholder  string `js:"placeholder"`
	Type         string `js:"type"`
	Value        string `js:"value"`
	DefaultValue string `js:"defaultValue"`
}

// InputProps creates a new instance of <input> properties, mutating the props
// by the provided initiainputser
func InputProps(f func(p *InputPropsDef)) *InputPropsDef {
	res := &InputPropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *InputDef) reactElement() {}

// Input creates a new instance of a <input> element with the provided props
func Input(props *InputPropsDef) *InputDef {
	args := []interface{}{"input", props}

	underlying := react.Call("createElement", args...)

	return &InputDef{
		underlying: underlying,
	}
}
