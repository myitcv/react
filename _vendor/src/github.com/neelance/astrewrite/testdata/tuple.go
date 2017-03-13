package main

func main() {
	g(f())
}

func f() (int, int) {
	return 0, 0
}

func g(x, y int) {
}
