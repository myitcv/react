package main

func test1() {
	_2 := makeChan()
	for {
		_, _1 := <-_2
		if !_1 {
			break
		}
		_3 := makeFunc()
		_3()
	}
}

func test2() {
	_2 := makeChan()
	for {
		x, _1 := <-_2
		if !_1 {
			break
		}
		_ = x
		_3 := makeFunc()
		_3()
	}
}

func test3() {
	var x int
	var _1 bool
	_2 := makeChan()
	for {
		x, _1 = <-_2
		if !_1 {
			break
		}
		_ = x
		_3 := makeFunc()
		_3()
	}
}

func makeChan() <-chan int {
	return nil
}

func makeFunc() func() int {
	return nil
}
