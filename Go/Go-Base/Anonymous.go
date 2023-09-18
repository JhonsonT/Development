package main

func foo() (int8, string) {
	return 100, "hello"
}

func main() {
	x, _ := foo()
	_, y := foo()
	println(x)
	println(y)

	const (
		a = iota
		_
		b
		c
		d
		e
	)

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
}
