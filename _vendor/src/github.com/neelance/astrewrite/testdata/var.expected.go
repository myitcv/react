package main

var x = func() { _1 := f(); _1() }

func f() func() {
	return nil
}
