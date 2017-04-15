// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package example // import "myitcv.io/immutable/example"

// The following directive will result in a generated file that includes the
// directive //go:generate echo "hello world"
//
//go:generate immutableGen -licenseFile license_header.txt -G "echo \"hello world\""

// MyMap will be exported
type _Imm_MyMap map[string]*MySlice

// MySlice will be exported
type _Imm_MySlice []*MyMap

// Where is this

// MyStruct will be exported.
//
// It is a special type.
type _Imm_MyStruct struct {
	// This is a non-field comment

	// Name is a field in MyStruct
	Name string `tag:"value"`

	// surname will not be exported
	surname string

	self *MyStruct

	// age will not be exported
	age int `tag:"age"`
}
