package markdowneditor

import (
	. "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/remarkable"
	"honnef.co/go/js/dom"
)

type MarkdownEditorDef struct {
	ComponentDef

	remark *remarkable.Remarkable
}

type MarkdownEditorState struct {
	value string
}

func MarkdownEditor() *MarkdownEditorDef {
	res := &MarkdownEditorDef{}
	res.remark = remarkable.NewRemarkable()

	BlessElement(res, nil)

	return res
}

func (p *MarkdownEditorDef) GetInitialState() MarkdownEditorState {
	return MarkdownEditorState{
		value: "Type some *markdown* here!",
	}
}

func (p *MarkdownEditorDef) Render() Element {
	return Div(nil,
		H3(nil, S("Input")),
		TextArea(
			TextAreaProps(func(tap *TextAreaPropsDef) {
				tap.ClassName = "form-control"
				tap.Value = p.State().value
				tap.OnChange = p.handleChange
			}),
		),
		H3(nil, S("Output")),
		Div(
			DivProps(func(dp *DivPropsDef) {
				dp.ClassName = "well"
				dp.DangerouslySetInnerHTML = p.getRawMarkup()
			}),
		),
	)
}

func (p *MarkdownEditorDef) handleChange(se *SyntheticEvent) {
	target := se.Target().(*dom.HTMLTextAreaElement)

	p.SetState(MarkdownEditorState{value: target.Value})
}

func (p *MarkdownEditorDef) getRawMarkup() *DangerousInnerHTMLDef {
	r := p.remark.Render(p.State().value)

	return DangerousInnerHTML(r)
}
