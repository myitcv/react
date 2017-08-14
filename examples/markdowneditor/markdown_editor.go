package markdowneditor // import "myitcv.io/react/examples/markdowneditor"

import (
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
func (m MarkdownEditorDef) Render() react.Element {
	val := m.State().value

	return react.Div(nil,
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
