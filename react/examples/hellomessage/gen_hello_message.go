package hellomessage

// Generated code (or at least will be once I write the code generator)
// for HelloMessage

func (p *HelloMessageDef) Props() HelloMessageProps {
	uprops := p.ComponentDef.Props()
	return uprops.(HelloMessageProps)
}
