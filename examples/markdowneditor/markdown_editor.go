package markdowneditor // import "myitcv.io/react/examples/markdowneditor"

import (
	"honnef.co/go/js/dom"
	r "myitcv.io/react"
	"myitcv.io/remarkable"
)

//go:generate reactGen

// MarkdownEditorDef is the definition of the MarkdownEditor component
type MarkdownEditorDef struct {
	r.ComponentDef
}

// MarkdownEditorState is the state type for the MarkdownEditor component
type MarkdownEditorState struct {
	value  string
	remark *remarkable.Remarkable
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
	}
}

// Render renders the MarkdownEditor component
func (m MarkdownEditorDef) Render() r.Element {
	val := m.State().value

	return r.Div(nil,
		r.H3(nil, r.S("Input")),
		r.TextArea(
			&r.TextAreaProps{
				ClassName: "form-control",
				Value:     val,
				OnChange:  inputChange{m},
			},
		),
		r.H3(nil, r.S("Output")),
		r.Div(
			&r.DivProps{
				ClassName:               "well",
				DangerouslySetInnerHTML: m.getRawMarkup(),
			},
		),
	)
}

func (m MarkdownEditorDef) getRawMarkup() *r.DangerousInnerHTML {
	st := m.State()
	md := st.remark.Render(st.value)
	return r.NewDangerousInnerHTML(md)
}

type inputChange struct{ m MarkdownEditorDef }

func (i inputChange) OnChange(se *r.SyntheticEvent) {
	st := i.m.State()

	target := se.Target().(*dom.HTMLTextAreaElement)
	st.value = target.Value

	i.m.SetState(st)
}
