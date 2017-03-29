package coretest

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
}

type _Imm_AS []*A

type _Imm_AM map[*A]*A

func main() {
}
