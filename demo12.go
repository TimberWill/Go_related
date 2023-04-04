package main

import "fmt"

func main() {
	a := 3
	b := 6.0

	c := int(b)
	d := float64(a)

	fmt.Printf("%T,%T,%T,%T", a, b, c, d)
}
