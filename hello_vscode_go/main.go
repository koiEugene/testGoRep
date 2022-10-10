package main

import "fmt"

/*
	func init() {
		fmt.Println("The first line.")
	}
*/
func main() {

	/*
		var var_1 bool
		var var_2 string
	*/
	var var_3, var_3_2 int
	/*var var_4 float32*/

	var_3 = 10
	var_3_2 = var_3

	fmt.Println("var_3 = ", var_3)
	fmt.Println("var_3_2 = ", var_3_2)

	var_3_2 += var_3_2
	var_3_2 /= 10
	fmt.Println("~var_3_2 = ", var_3_2)

	if var_3 == var_3_2 {
		fmt.Println("условный оператор = ", var_3_2)
	} else {
		fmt.Println("условный оператор = ", var_3)
	}

	/*
		fmt.Println(var_1, var_2, var_3, var_4)

		message := "Hello World. IT's Golang!"
		fmt.Println(message)
	*/
}
