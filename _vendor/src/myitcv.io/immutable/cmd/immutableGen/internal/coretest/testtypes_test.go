package coretest

// a comment about MyMap
type _Imm_MyTestMap map[string]int

// a comment about Slice
type _Imm_MyTestSlice []*string

// a comment about myStruct
type _Imm_MyTestStruct struct {

	// my field comment
	//somethingspecial
	/*

		Heelo

	*/
	Name, surname string `tag:"value"`
	age           int    `tag:"age"`

	fieldWithoutTag bool
}
