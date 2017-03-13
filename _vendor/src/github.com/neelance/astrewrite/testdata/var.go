package main

var x = func() { f()() }

func f() func() {
	return nil
}
