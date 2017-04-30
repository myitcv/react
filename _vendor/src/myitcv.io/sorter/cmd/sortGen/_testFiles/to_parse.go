package main

//go:generate sortGen
//go:generate immutableGen

import (
	"bytes"
	"fmt"

	"myitcv.io/sorter"
	"myitcv.io/sorter/cmd/sortGen/_testFiles/internal/other"
)

type _Imm_MySlice []string

type person struct {
	name string
	age  int
}

func main() {
	people := []person{
		person{"Sarah", 60},
		person{"Jill", 34},
		person{"Paul", 25},
	}

	fmt.Printf("Before: %v\n", people)

	sortByName(people)

	fmt.Printf("Name sorted: %v\n", people)

	SortByAge(people)

	fmt.Printf("Age sorted: %v\n", people)
}

// MATCH
func orderByName(persons []person, i, j int) sorter.Ordered {
	return persons[i].name < persons[j].name
}

// MATCH - same package
func orderMySlice(things *MySlice, i, j int) sorter.Ordered {
	return things.Get(i) < things.Get(j)
}

// MATCH - other package
func orderOtherMySlice(things *other.MySlice, i, j int) sorter.Ordered {
	return things.Get(i) < things.Get(j)
}

// fail
func order(persons []person, i, j int) sorter.Ordered {
	return persons[i].name < persons[j].name
}

// MATCH
func orderPointerByName(persons []*person, i, j int) sorter.Ordered {
	return persons[i].name < persons[j].name
}

// MATCH
func orderBufferByContents(buffers []bytes.Buffer, i int, j int) sorter.Ordered {
	return buffers[i].String() < buffers[j].String()
}

// MATCH
func orderMap(buffers []map[string]bool, i int, j int) sorter.Ordered {
	return true
}

type example struct{}

// MATCH
func (e *example) orderBanana(s []string, i, j int) sorter.Ordered {
	return s[i] < s[j]
}
