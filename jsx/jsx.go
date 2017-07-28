package jsx

import (
	"fmt"
	"strings"

	"github.com/russross/blackfriday"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	r "myitcv.io/react"
	rhtml "myitcv.io/react/html"
)

// each of the parse* functions does zero validation
// this is intentional because everything is expected to
// go via the generic parse function

// TODO code generate these parse functions

func parseP(n *html.Node) *rhtml.PElem {
	var kids []r.Element

	// TODO attributes

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.P(nil, kids...)
}

func parseHR(n *html.Node) *rhtml.HrElem {
	// TODO attributes

	return rhtml.Hr(nil)
}

func parseBR(n *html.Node) *rhtml.BrElem {
	// TODO attributes

	return rhtml.Br(nil)
}

func parseH1(n *html.Node) *rhtml.H1Elem {
	var kids []r.Element

	// TODO attributes

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.H1(nil, kids...)
}

func parseSpan(n *html.Node) *rhtml.SpanElem {
	var kids []r.Element

	var vp *rhtml.SpanProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.SpanProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "classname":
				vp.ClassName = a.Val
			case "style":
				vp.Style = parseCSS(a.Val)
			default:
				panic(fmt.Errorf("don't know how to handle <span> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Span(vp, kids...)
}

func parseI(n *html.Node) *rhtml.IElem {
	var kids []r.Element

	var vp *rhtml.IProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.IProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "id":
				vp.ID = a.Val
			case "classname":
				vp.ClassName = a.Val
			default:
				panic(fmt.Errorf("don't know how to handle <i> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.I(vp, kids...)
}

func parseFooter(n *html.Node) *rhtml.FooterElem {
	var kids []r.Element

	var vp *rhtml.FooterProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.FooterProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "id":
				vp.ID = a.Val
			case "classname":
				vp.ClassName = a.Val
			default:
				panic(fmt.Errorf("don't know how to handle <footer> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Footer(vp, kids...)
}

func parseDiv(n *html.Node) *rhtml.DivElem {
	var kids []r.Element

	var vp *rhtml.DivProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.DivProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "id":
				vp.ID = a.Val
			case "classname":
				vp.ClassName = a.Val
			case "style":
				vp.Style = parseCSS(a.Val)
			default:
				panic(fmt.Errorf("don't know how to handle <div> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Div(vp, kids...)
}

func parseButton(n *html.Node) *rhtml.ButtonElem {
	var kids []r.Element

	var vp *rhtml.ButtonProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.ButtonProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "id":
				vp.ID = a.Val
			case "classname":
				vp.ClassName = a.Val
			default:
				panic(fmt.Errorf("don't know how to handle <div> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Button(vp, kids...)
}

func parseCode(n *html.Node) *rhtml.CodeElem {
	var kids []r.Element

	// TODO attributes

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Code(nil, kids...)
}

func parseH3(n *html.Node) *rhtml.H3Elem {
	var kids []r.Element

	// TODO attributes

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.H3(nil, kids...)
}

func parseImg(n *html.Node) *rhtml.ImgElem {
	var kids []r.Element

	var vp *rhtml.ImgProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.ImgProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "src":
				vp.Src = a.Val
			case "style":
				vp.Style = parseCSS(a.Val)
			default:
				panic(fmt.Errorf("don't know how to handle <img> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.Img(vp, kids...)
}

func parseA(n *html.Node) *rhtml.AElem {
	var kids []r.Element

	var vp *rhtml.AProps

	if len(n.Attr) > 0 {
		vp = new(rhtml.AProps)

		for _, a := range n.Attr {
			switch a.Key {
			case "href":
				vp.Href = a.Val
			case "target":
				vp.Target = a.Val
			default:
				panic(fmt.Errorf("don't know how to handle <a> attribute %q", a.Key))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		kids = append(kids, parse(c))
	}

	return rhtml.A(vp, kids...)
}

// TODO replace with proper parser
func parseCSS(s string) *rhtml.CSS {
	res := new(rhtml.CSS)

	parts := strings.Split(s, ";")

	for _, p := range parts {
		kv := strings.Split(p, ":")
		if len(kv) != 2 {
			panic(fmt.Errorf("invalid key-val %q in %q", p, s))
		}

		k, v := kv[0], kv[1]

		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\"")

		switch k {
		case "overflow-y":
			res.OverflowY = v
		case "margin-top":
			res.MarginTop = v
		case "font-size":
			res.FontSize = v
		case "font-style":
			res.FontStyle = v
		default:
			panic(fmt.Errorf("unknown CSS key %q in %q", k, s))
		}
	}

	return res
}

func parse(n *html.Node) r.Element {
	switch n.Type {
	case html.TextNode:
		return r.S(n.Data)
	case html.ElementNode:
		// we will fall out from here...
	default:
		panic(fmt.Errorf("cannot handle NodeType %v", n.Type))
	}

	switch n.Data {
	case "p":
		return parseP(n)
	case "h1":
		return parseH1(n)
	case "code":
		return parseCode(n)
	case "h3":
		return parseH3(n)
	case "img":
		return parseImg(n)
	case "a":
		return parseA(n)
	case "footer":
		return parseFooter(n)
	case "div":
		return parseDiv(n)
	case "span":
		return parseSpan(n)
	case "hr":
		return parseHR(n)
	case "br":
		return parseBR(n)
	case "button":
		return parseButton(n)
	case "i":
		return parseI(n)
	default:
		panic(fmt.Errorf("cannot handle Element %v", n.Data))
	}
}

var htmlCache = make(map[string][]r.Element)

// HTML is a runtime JSX-like parserhtml. It parses the supplied HTML string into
// myitcv.io/react element values. It exists as a stop-gap runtime solution to
// full JSX-like support within the GopherJS compilerhtml. It should only be used
// where the argument is a compile-time constant string (TODO enforce this
// within reactVet). HTML will panic in case s cannot be parsed as a valid HTML
// fragment
//
func HTML(s string) []r.Element {
	s = strings.TrimSpace(s)

	if v, ok := htmlCache[s]; ok {
		return v
	}

	// a dummy div for parsing the fragment
	div := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}

	elems, err := html.ParseFragment(strings.NewReader(s), div)
	if err != nil {
		panic(fmt.Errorf("failed to parse HTML %q: %v", s, err))
	}

	res := make([]r.Element, len(elems))

	for i, v := range elems {
		res[i] = parse(v)
	}

	htmlCache[s] = res

	return res
}

// HTMLElem is a convenience wrapper around HTML where only a single root
// element is expected. HTMLElem will panic if more than one HTML element
// results
//
func HTMLElem(s string) r.Element {
	res := HTML(s)

	if v := len(res); v != 1 {
		panic(fmt.Errorf("expected single element result from %q; got %v", s, v))
	}

	return res[0]
}

// Markdown is a runtime JSX-like parser for markdown. It parses the supplied
// markdown string into an HTML string and then hands off to the HTML function.
// Like the HTML function, it exists as a stop-gap runtime solution to full
// JSX-like support within the GopherJS compilerhtml. It should only be used where
// the argument is a compile-time constant string (TODO enforce this within
// reactVet). Markdown will panic in case the markdown string s results in an
// invalid HTML string
//
func Markdown(s string) []r.Element {

	h := blackfriday.MarkdownCommon([]byte(s))

	return HTML(string(h))
}
