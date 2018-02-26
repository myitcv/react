package example_test

import (
	"fmt"
	"testing"

	"myitcv.io/immutable/example"
)

func TestThis(t *testing.T) {
	// the zero value of Person is immutable
	p1 := new(example.Person)

	fmt.Printf("p1: %v\n", p1)
	fmt.Println()

	// hence setting the name on the Person pointed to by p1 leaves that Person
	// unchanged and instead returns a new Person with the name set (notice the
	// code generated "setter" for Name)
	p2 := p1.SetName("Paul")

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Println()

	p4 := p2.SetAge(42)

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p4: %v\n", p4)
	fmt.Println()

	p3 := p1.WithMutable(func(p *example.Person) {
		// WithMutable is used whre multiple mutations are required... but again
		// the mutations are applied to a copy (the p passed to this function is
		// mutable copy of p1, which when WithMutable returns is marked
		// immutable)
		p.SetName("Monty Python")
		p.SetAge(42)
	})

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p3: %v\n", p3)
}
