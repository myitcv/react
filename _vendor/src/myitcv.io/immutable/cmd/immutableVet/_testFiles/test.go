package _testFiles

//go:generate immutableGen

type _Imm_Dummy struct {
	Name string
}

type _Imm_Dummy2 struct {
	name    []byte // ERROR
	other   *Dummy3
	mine    MyIntf // ERROR
	another MyType // ERROR
}

type _Imm_Dummy3 struct {
	other *Dummy2
}

type another Dummy // ERROR

type _Imm_intS []int

var _ _Imm_Dummy // ERROR
var _ Dummy      // ERROR

var _ map[Dummy]string // ERROR
var _ map[string]Dummy // ERROR
var _ []Dummy          // ERROR

var _ = func(d Dummy) Dummy { // ERROR x 2
	return Dummy{} // ERROR
}

func Eg(d Dummy) Dummy { // ERROR x 2
	return Dummy{} // ERROR
}

var _ = &Dummy{} // ERROR

var _ = new(Dummy)

var good *Dummy

func main() {
	print(good._Name) // ERROR

	var x *intS

	_ = x.Range() // ERROR

	for _ = range x.Range() {
	}

	y := *x
	print(y)

	_ = append([]int{}, x.Range()...)

	x.Append(x.Range()...)
}

type MyIntf interface {
	IsMyIntf()
}

type MyType []string

func (m MyType) IsMyIntf() {}

var _ MyIntf = MyType([]string{})

var _ = Dummy{} // ERROR
