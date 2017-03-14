package markdowneditor

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/remarkable"
	"honnef.co/go/js/dom"
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
			r.TextAreaProps(func(tap *r.TextAreaPropsDef) {
				tap.ClassName = "form-control"
				tap.Value = m.State().value
				tap.OnChange = m.handleChange
			}),
		),
		r.H3(nil, r.S("Output")),
		r.Div(
			r.DivProps(func(dp *r.DivPropsDef) {
				dp.ClassName = "well"
				dp.DangerouslySetInnerHTML = m.getRawMarkup()
			}),
		),
	)
}

func (m *MarkdownEditorDef) handleChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLTextAreaElement)

	m.SetState(MarkdownEditorState{value: target.Value})
}

func (m *MarkdownEditorDef) getRawMarkup() *r.DangerousInnerHTMLDef {
	rem := m.remark.Render(m.State().value)

	return r.DangerousInnerHTML(rem)
}
