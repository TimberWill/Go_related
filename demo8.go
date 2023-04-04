package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "ok"
		e
		f = 100
		g = iota
	)
	const (
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
