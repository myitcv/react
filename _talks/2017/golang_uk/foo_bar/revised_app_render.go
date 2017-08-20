// app.go

func (a AppDef) Render() react.Element {
	return react.Div(nil,
		react.H1(nil, react.S("Hello World")),
		react.P(nil, react.S("This is my first GopherJS React App.")),

		FooBar( // HL
			FooBarProps{ // HL
				Name: "Peter", // HL
			}, // HL
		), // HL
	)
}

