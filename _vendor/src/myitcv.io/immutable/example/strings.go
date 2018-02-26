package example

import "fmt"

//go:generate immutableGen

// via go generate, this template is code generated into the immutable Person
// struct within the same package
type _Imm_Person struct {
	Name string
	Age  int
}

// Hence we can then define methods on *Person (methods can only be defined on
// a pointer receiver)
func (p *Person) String() string {
	return fmt.Sprintf("Person{ Name: %q, Age: %v}", p.Name(), p.Age())
}
