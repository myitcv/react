package markdowneditor // import "myitcv.io/react/examples/markdowneditor"

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"myitcv.io/react"
	"myitcv.io/remarkable"
)

//go:generate reactGen

// MarkdownEditorDef is the definition of the MarkdownEditor component
type MarkdownEditorDef struct {
	react.ComponentDef
}

// MarkdownEditorState is the state type for the MarkdownEditor component
type MarkdownEditorState struct {
	value  string
	remark *remarkable.Remarkable

	// We don't actually use the DOM element for the containing div in the
	// logic of the MarkdownEditor example, rather it's just a demonstration
	// of React Refs at work.
	div *divRef
}

// MarkdownEditor creates instances of the MarkdownEditor component
func MarkdownEditor() *MarkdownEditorElem {
	return buildMarkdownEditorElem()
}

// GetInitialState returns the initial state for a MarkdownEditor component
func (m MarkdownEditorDef) GetInitialState() MarkdownEditorState {
	remark := remarkable.NewRemarkable()
	return MarkdownEditorState{
		value:  "Type some *markdown* here!",
		remark: remark,
		div:    &divRef{m: m},
	}
}

// Render renders the MarkdownEditor component
func (m MarkdownEditorDef) Render() react.Element {
	val := m.State().value

	return react.Div(&react.DivProps{Ref: m.State().div},
		react.H3(nil, react.S("Input")),
		react.TextArea(
			&react.TextAreaProps{
				ClassName: "form-control",
				Value:     val,
				OnChange:  inputChange{m},
			},
		),
		react.H3(nil, react.S("Output")),
		react.Div(
			&react.DivProps{
				ClassName:               "well",
				DangerouslySetInnerHTML: m.getRawMarkup(),
			},
		),
	)
}

func (m MarkdownEditorDef) getRawMarkup() *react.DangerousInnerHTML {
	st := m.State()
	md := st.remark.Render(st.value)
	return react.NewDangerousInnerHTML(md)
}

type inputChange struct{ m MarkdownEditorDef }

func (i inputChange) OnChange(se *react.SyntheticEvent) {
	st := i.m.State()

	target := se.Target().(*dom.HTMLTextAreaElement)
	st.value = target.Value

	i.m.SetState(st)
}

type divRef struct {
	m   MarkdownEditorDef
	div *dom.HTMLDivElement
}

func (d *divRef) Ref(h *js.Object) {
	var div *dom.HTMLDivElement
	if e := dom.WrapHTMLElement(h); e != nil {
		div = e.(*dom.HTMLDivElement)
	}

	d.div = div

	print("Here is the containing div for the rendered MarkdownEditor", div.Object)

	// in case we need to re-render at this point we could call
	d.m.ForceUpdate()
}
