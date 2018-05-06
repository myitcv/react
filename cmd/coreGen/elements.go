package main

import "fmt"

type Elem struct {
	// The myitcv.io/react Name of the element - not set directly, taken from
	// the key of the elements map.
	Name string

	// React is an override for the React name of the element if it is otherwise
	// not equal to the lowercase version of .Name
	React string

	// Dom is the name used by honnef.co/go/js/dom when referring to the underlying
	// HTML element. Default is .Name
	Dom string

	// Attributes maps the name of an attribute to the definition of an
	// attribute.
	Attributes map[string]*Attr

	// NonBasic is true if  honnef.co/go/js/dom does not declare a specific
	// Element type.
	NonBasic bool

	// Templates lists the attribute templates this element should use as a
	// base.
	Templates []string

	// NonHTML indicates this element should not automatically inherit the html
	// attribute template
	NonHTML bool

	// Child indicates this element can take a single child of the provided type.
	// Its use is exclusive with Children. No default value.
	Child string

	// Children indicates this element can take a multiple children of the provided
	// type. Its use is exclusive with Child. Default is Element.
	Children string

	// EmptyElement indicates the element may not have any children
	EmptyElement bool

	// Implements is the list of special interface methods this element implements.
	Implements []string

	// SkipTests is an override on whether to not generate the boilerplate tests.
	SkipTests bool
}

func (e *Elem) ChildParam() string {
	if e.Child != "" {
		return "child " + e.Child
	} else if e.Children != "" {
		return "children ..." + e.Children
	}

	return ""
}

func (e *Elem) ChildConvert() string {
	if e.Children != "" && e.Children != "Element" {
		return `
var elems []Element
for _, v := range children {
	elems = append(elems, v)
}
		`
	}

	return ""
}

func (e *Elem) ChildArg() string {
	if e.Child != "" {
		return "child"
	} else if e.Children != "" {
		if e.Children == "Element" {
			return "children..."
		} else {
			return "elems..."
		}
	}

	return ""
}

func (e *Elem) ChildrenReactType() string {
	if e.Children[0] == '*' {
		return "*react." + e.Children[1:]
	}

	return "react." + e.Children
}

func (e *Elem) HTMLAttributes() map[string]*Attr {
	res := make(map[string]*Attr)

	for n, a := range e.Attributes {
		if a.NoHTML || a.NoReact || a.IsEvent || a.Name == "Ref" {
			continue
		}

		res[n] = a
	}

	return res
}

type Attr struct {
	// The myitcv.io/react Name of the attribute - not set directly, taken from
	// the key of the elements map.
	Name string

	// React is an override for the React name of the attribute if it is otherwise
	// not equal to the lower-initial version of .Name
	React string

	// HTML is an override for the HTML attribute name if it is otherwise not equal
	// to the lowercase version of .Name
	HTML string

	// HTMLConvert is a function that must be called on a JSX-parsed value before
	// assignment. Default is nothing.
	HTMLConvert string

	// Type is an override for the type of the attribute. The zero value implies
	// string
	Type string

	// OmitEmpty indicates that no attribute should be set on the underlying React
	// element if the zero value of the attribute is set.
	OmitEmpty bool

	// NoReact indicates that this attribute should not attempt to be mapped directly
	// to an underlying React attribute.
	NoReact bool

	// NoHTML indicates this attribute does not have an HTML equivalent, and hence
	// should not appear during parsing.
	NoHTML bool

	// IsEvent indicates that the attribute is an event.
	IsEvent bool
}

func (a *Attr) Tag() string {
	omitEmpty := ""
	if a.OmitEmpty {
		omitEmpty = ` react:"omitempty"`
	}
	return fmt.Sprintf("`js:\"%v\"%v`", a.React, omitEmpty)
}

func (a *Attr) HTMLConvertor(s string) string {
	if a.HTMLConvert == "" {
		return s
	}

	return fmt.Sprintf("%v(%v)", a.HTMLConvert, s)
}

// templates are the attribute templates to which elements can refer
var templates = map[string]map[string]*Attr{
	"html": {
		"AriaHasPopup":            &Attr{React: "aria-haspopup", Type: "bool", HTML: "aria-haspopup"},
		"AriaExpanded":            &Attr{React: "aria-expanded", Type: "bool", HTML: "aria-expanded"},
		"AriaLabelledBy":          &Attr{React: "aria-labelledby", HTML: "aria-labelledby"},
		"ClassName":               &Attr{HTML: "class"},
		"DangerouslySetInnerHTML": &Attr{Type: "*DangerousInnerHTML", NoHTML: true},
		"DataSet":                 &Attr{Type: "DataSet", NoReact: true},
		"ID":                      &Attr{OmitEmpty: true, React: "id"},
		"Key":                     &Attr{OmitEmpty: true},
		"Ref":                     &Attr{Type: "Ref"},
		"Role":                    &Attr{},
		"Style":                   &Attr{Type: "*CSS", HTMLConvert: "parseCSS"},

		// Events
		"OnChange": &Attr{Type: "OnChange", IsEvent: true},
		"OnClick":  &Attr{Type: "OnClick", IsEvent: true},
	},
}

// elements is a map from the Go element name to the definition
var elements = map[string]*Elem{
	"A": &Elem{
		Dom: "HTMLAnchorElement",
		Attributes: map[string]*Attr{
			"Href":   &Attr{},
			"Target": &Attr{},
			"Title":  &Attr{},
		},
	},
	"Div": &Elem{},
	"H1": &Elem{
		Dom: "HTMLHeadingElement",
	},
	"H2": &Elem{
		Dom: "HTMLHeadingElement",
	},
	"H3": &Elem{
		Dom: "HTMLHeadingElement",
	},
	"H4": &Elem{
		Dom: "HTMLHeadingElement",
	},
	"H5": &Elem{
		Dom: "HTMLHeadingElement",
	},
	"Br": &Elem{
		Dom: "HTMLBRElement",
	},
	"TextArea": &Elem{
		Attributes: map[string]*Attr{
			"Placeholder": &Attr{},
			"Value":       &Attr{},
		},
	},
	"Button": &Elem{
		Attributes: map[string]*Attr{
			"Type": &Attr{},
		},
	},
	"Ul": &Elem{
		Dom:      "HTMLUListElement",
		Children: "RendersLi",
	},
	"Li": &Elem{
		Dom:        "HTMLLIElement",
		Implements: []string{"RendersLi(*LiElem)"},
	},
	"Span": &Elem{},
	"Pre":  &Elem{},
	"Nav": &Elem{
		Dom: "BasicHTMLElement",
	},
	"Code": &Elem{
		Dom: "BasicHTMLElement",
	},
	"IFrame": &Elem{
		Attributes: map[string]*Attr{
			"SrcDoc": &Attr{},
		},
	},
	"Select": &Elem{
		Attributes: map[string]*Attr{
			"Value": &Attr{},
		},
		Children: "*OptionElem",
	},
	"Option": &Elem{
		Attributes: map[string]*Attr{
			"Value": &Attr{},
		},
	},
	"Img": &Elem{
		Dom: "HTMLImageElement",
		Attributes: map[string]*Attr{
			"Src": &Attr{},
			"Alt": &Attr{},
		},
	},
	"Form": &Elem{},
	"Label": &Elem{
		Attributes: map[string]*Attr{
			"For": &Attr{
				React: "htmlFor",
			},
		},
	},
	"Strike": &Elem{
		Dom: "BasicHTMLElement",
	},
	"P": &Elem{
		Dom: "HTMLParagraphElement",
	},
	"I": &Elem{
		Dom: "BasicHTMLElement",
	},
	"Table": &Elem{},
	"Footer": &Elem{
		Dom: "BasicHTMLElement",
	},
	"Hr": &Elem{
		Dom:          "HTMLHRElement",
		EmptyElement: true,
	},
	"Input": &Elem{
		Attributes: map[string]*Attr{
			"Placeholder": &Attr{},
			"Type":        &Attr{},
			"Value":       &Attr{},
		},
	},
}
