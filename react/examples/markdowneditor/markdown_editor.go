package markdowneditor

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/remarkable"
	"honnef.co/go/js/dom"
)

//go:generate reactGen

type MarkdownEditorDef struct {
	r.ComponentDef

	remark *remarkable.Remarkable
}

type MarkdownEditorState struct {
	value string
}

func MarkdownEditor() *MarkdownEditorDef {
	res := &MarkdownEditorDef{}
	res.remark = remarkable.NewRemarkable()

	r.BlessElement(res, nil)

	return res
}

func (p *MarkdownEditorDef) GetInitialState() MarkdownEditorState {
	return MarkdownEditorState{
		value: "Type some *markdown* here!",
	}
}

func (p *MarkdownEditorDef) Render() r.Element {
	return r.Div(nil,
		r.H3(nil, r.S("Input")),
		r.TextArea(
			r.TextAreaProps(func(tap *r.TextAreaPropsDef) {
				tap.ClassName = "form-control"
				tap.Value = p.State().value
				tap.OnChange = p.handleChange
			}),
		),
		r.H3(nil, r.S("Output")),
		r.Div(
			r.DivProps(func(dp *r.DivPropsDef) {
				dp.ClassName = "well"
				dp.DangerouslySetInnerHTML = p.getRawMarkup()
			}),
		),
	)
}

func (p *MarkdownEditorDef) handleChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLTextAreaElement)

	p.SetState(MarkdownEditorState{value: target.Value})
}

func (p *MarkdownEditorDef) getRawMarkup() *r.DangerousInnerHTMLDef {
	rem := p.remark.Render(p.State().value)

	return r.DangerousInnerHTML(rem)
}
