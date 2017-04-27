## `sortGen`

A [`go generate`](https://blog.golang.org/generate)-or that makes sorting arbitrary slices easier via order functions, removing
the need to implement [`sort.Interface`](https://godoc.org/sort#Interface) etc.

Simply define an order function on the slice type of interest:

```go
// main.go
//go:generate sortGen

func orderByName(persons []person, i, j int) sorter.Ordered {
	return persons[i].name < persons[j].name
}
```

then run `go generate` and corresponding sort functions will have been generated for you:

```go
// gen_main_sorter.go

func sortByName(vs []person) {
	...
}

func stableSortByName(vs []person) {
	...
}
```

See the example and rules below for more details.

### Install

```
go get -u myitcv.io/sorter/cmd/sortGen
```

### Example

Taking the example from [`example/main.go`](https://myitcv.io/sorter/blob/master/example/main.go):

```go
//go:generate sortGen -licenseFile license_header.txt

package main

import (
	"fmt"

	"myitcv.io/sorter"
)

type person struct {
	name string
	age  int
}

func main() {
	people := []person{
		person{"Sarah", 25},
		person{"Jill", 34},
		person{"Paul", 25},
	}

	fmt.Printf("Before: %v\n", people)

	sortByName(people)

	fmt.Printf("Name sorted: %v\n", people)

	m := &myStruct{}
	m.stableSortByAge(people)

	fmt.Printf("Age sorted: %v\n", people)
}

func orderByName(persons []person, i, j int) sorter.Ordered {
	return persons[i].name < persons[j].name
}

type myStruct struct {
	// some fields
}

func (m *myStruct) orderByAge(persons []person, i, j int) sorter.Ordered {
	return persons[i].age < persons[j].age
}
```

Then:

```
$ go generate
$ go build
$ ./example
Before: [{Sarah 25} {Jill 34} {Paul 25}]
Name sorted: [{Jill 34} {Paul 25} {Sarah 25}]
Age sorted: [{Paul 25} {Sarah 25} {Jill 34}]
```

Examine the contents of [`gen_main_sorter.go`](https://myitcv.io/sorter/blob/master/example/gen_main_sorter.go) to see the generated functions.

### Features

* Supports `-licenseFile FILENAME` flag which allows a file containing an uncommented license header
to be included (commented) at the top of each generated file

### Rules

`sortGen` generates sort and stable sort functions/methods according to the following simple rules:

1. The file, e.g. `my_file.go`, containing the order function/method must include the directive `//go:generate sortGen`
2. The order function/method name must be of the form `"order*"` or `"Order*"` (more strictly `^[oO]rder[[:word:]]+` in a [regex](https://godoc.org/regexp)
   [pattern](https://github.com/google/re2/wiki/Syntax))
3. The parameters of the order function/method must be a slice type, followed by two `int`'s
4. The return type must be `myitcv.io/sorter.Ordered`

The sort functions/methods generated will be of the form `"sort*"` or `"Sort*"` and `"stableSort*"` or `"StableSort*"`
(following the capitalisation of the order function). They will be written to a file with a name corresponding to the
input file, `gen_my_file_sorter.go` in the case of the file mentioned in point 1.

Notice the use of the term "function/method"; if the order function defines a receiver (i.e. it is a method)
then the generated sort (and stable sort) function will use the same receiver, and hence be a method.

### Implementation

The current implementation of the generator simply wraps a call to `sort.Sort` (or `sort.Stable`); this of course can be improved...

### Bugs

This is only an initial proof-of-concept, probably lots of bugs and edge cases missed. Please raise issues...
