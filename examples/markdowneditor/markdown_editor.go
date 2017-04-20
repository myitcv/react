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

	remark *remarkable.Remarkable
}

// MarkdownEditorState is the state type for the MarkdownEditor component
type MarkdownEditorState struct {
	value string
}

// MarkdownEditor creates instances of the MarkdownEditor component
func MarkdownEditor() *MarkdownEditorDef {
	res := &MarkdownEditorDef{}
	res.remark = remarkable.NewRemarkable()

	r.BlessElement(res, nil)

	return res
}

// GetInitialState returns the initial state for a MarkdownEditor component
func (m *MarkdownEditorDef) GetInitialState() MarkdownEditorState {
	return MarkdownEditorState{
		value: "Type some *markdown* here!",
	}
}

// Render renders the MarkdownEditor component
func (m *MarkdownEditorDef) Render() r.Element {
	return r.Div(nil,
		r.H3(nil, r.S("Input")),
		r.TextArea(
			&r.TextAreaProps{
				ClassName: "form-control",
				Value:     m.State().value,
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

func (m *MarkdownEditorDef) getRawMarkup() *r.DangerousInnerHTMLDef {
	rem := m.remark.Render(m.State().value)

	return r.DangerousInnerHTML(rem)
}

type inputChange struct{ m *MarkdownEditorDef }

func (i inputChange) OnChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLTextAreaElement)

	i.m.SetState(MarkdownEditorState{value: target.Value})
}
