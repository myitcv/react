package react

import "github.com/gopherjs/gopherjs/js"

// H1Def is the React component definition corresponding to the HTML <h1> element
type H1Def struct {
	underlying *js.Object
}

// H1PropsDef defines the properties for the <h1> element
type H1PropsDef struct {
	*BasicHTMLElement
}

// H1Props creates a new instance of <h1> properties, mutating the props
// by the provided initiah1ser
func H1Props(f func(p *H1PropsDef)) *H1PropsDef {
	res := &H1PropsDef{
		BasicHTMLElement: newBasicHTMLElement(),
	}
	f(res)
	return res
}

func (d *H1Def) reactElement() {}

// H1 creates a new instance of a <h1> element with the provided props and
// child
func H1(props *H1PropsDef, child Element) *H1Def {
	underlying := react.Call("createElement", "h1", props, elementToReactObj(child))

	return &H1Def{underlying: underlying}
}
