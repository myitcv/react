package coretest

import "myitcv.io/immutable"

//go:generate immutableGen -licenseFile license.txt -G "echo \"hello world\""

// a comment about MyMap
type _Imm_MyMap map[string]int

// a comment about Slice
type _Imm_MySlice []string

type MyStructUuid uint64

type MyStructKey struct {
	Uuid    MyStructUuid
	Version uint64
}

// a comment about myStruct
type _Imm_MyStruct struct {
	Key MyStructKey

	// my field comment
	//somethingspecial
	/*

		Heelo

	*/
	Name, surname string `tag:"value"`
	age           int    `tag:"age"`

	string

	fieldWithoutTag bool
}

type _Imm_A struct {
	Name string
	A    *A

	Blah
}

type _Imm_AS []*A

type _Imm_AM map[*A]*A

type Blah interface {
	immutable.Immutable
}

type _Imm_BlahUse struct {
	Blah
}

type BlahMutable struct{}

var _ Blah = BlahMutable{}

func (b BlahMutable) Mutable() bool {
	return true
}

func (b BlahMutable) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	return false
}

type BlahNonMutable struct{}

var _ Blah = BlahNonMutable{}

func (b BlahNonMutable) Mutable() bool {
	return false
}

func (b BlahNonMutable) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	return true
}

func main() {
}
