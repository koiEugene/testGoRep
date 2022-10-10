package main

import "fmt"

func abc(x, y int) int {
	return x + y
}
func main() {
	f := abc(2, 3)
	fmt.Println("resalt = ", f)
	a := abc(3, 4)
	fmt.Println("result a =", a)
}
