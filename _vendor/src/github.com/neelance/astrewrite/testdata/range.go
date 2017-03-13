package main

func test1() {
	for range makeChan() {
		makeFunc()()
	}
}

func test2() {
	for x := range makeChan() {
		_ = x
		makeFunc()()
	}
}

func test3() {
	var x int
	for x = range makeChan() {
		_ = x
		makeFunc()()
	}
}

func makeChan() <-chan int {
	return nil
}

func makeFunc() func() int {
	return nil
}
