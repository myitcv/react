// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package sorter

// Inspired by https://github.com/mattn/sorter

// Ordered is a named type used to help identify order functions. See sortGen
// for more details
type Ordered bool

const (
	// OrderedName is the name of the named type above, exposed here so that
	// generators can safely refer to it without it breaking
	OrderedName = "Ordered"

	// PkgName is the package to which this definition belongs, again exposed
	// here for the benefit of generators
	PkgName = "myitcv.io/sorter"
)

// Wrapper is a light wrapper to faciliate calls to sort.Sort
type Wrapper struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (w *Wrapper) Len() int {
	return w.LenFunc()
}

func (w *Wrapper) Less(i, j int) bool {
	return w.LessFunc(i, j)
}

func (w *Wrapper) Swap(i, j int) {
	w.SwapFunc(i, j)
}
