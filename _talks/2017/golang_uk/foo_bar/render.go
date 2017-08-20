func (f FooBarDef) Render() react.Element {

	name := f.Props().Name // HL
	age := f.State().Age // HL

	details := fmt.Sprintf("My name is %v. My age is %v", name, age)

	return react.Div(nil,
		react.P(nil,
			react.S(details),
		),
		react.Button(
			&react.ButtonProps{
				OnClick: ageClick{f},
			},
			react.S("Bump age"),
		),
	)
}

