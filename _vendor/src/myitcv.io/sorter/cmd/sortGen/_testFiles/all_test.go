package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	// we skip any effort to work out whether the sort
	// is correct or not; and instead just verify
	// that we have generated the wrappers
	sortByName(nil)
	stableSortByName(nil)

	sortMap(nil)
	stableSortMap(nil)

	sortPointerByName(nil)
	stableSortPointerByName(nil)

	sortBufferByContents(nil)
	stableSortBufferByContents(nil)

	SortByAge(nil)
	StableSortByAge(nil)

	m1 := NewMySlice("banana", "apple")
	m2 := sortMySlice(m1)

	for i, v := range m1.Range() {
		fmt.Printf("%v: %v\n", i, v)
	}

	for i, v := range m2.Range() {
		fmt.Printf("%v: %v\n", i, v)
	}

	sortOtherMySlice(nil)
	stableSortOtherMySlice(nil)
}
