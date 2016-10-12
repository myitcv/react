package markdowneditor

import "github.com/myitcv/gopherjs/react"

// Generated code (or at least will be once I write the code generator)
// for MarkdownEditor

func (p *MarkdownEditorDef) SetState(s MarkdownEditorState) {
	p.ComponentDef.SetState(s)
}

func (p *MarkdownEditorDef) State() MarkdownEditorState {
	return p.ComponentDef.State().(MarkdownEditorState)
}

func (p *MarkdownEditorDef) GetInitialStateIntf() react.State {
	return p.GetInitialState()
}

func (p MarkdownEditorState) IsState() {}
