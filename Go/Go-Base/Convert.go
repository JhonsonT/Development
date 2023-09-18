package main

import "math"

func main() {
	var a, b = 2, 3
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	println(c)
}
