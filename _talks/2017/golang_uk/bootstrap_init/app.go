// AppDef is the definition of the App component
//
type AppDef struct { // HL
	react.ComponentDef
}

// App creates instances of the App component
//
func App() *AppElem {
	return buildAppElem()
}

// Render renders the App component
//
func (a AppDef) Render() react.Element { // HL
	return react.Div(nil,
		react.H1(nil, react.S("Hello World")),
		react.P(nil, react.S("This is my first GopherJS React App.")),
	)
}

