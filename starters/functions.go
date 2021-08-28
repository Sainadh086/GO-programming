// implementing functions
package main

import "fmt"

func main() {
	superman()
	result := multiple(3, 6)
	fmt.Println(result)
	mysum, adderLen, fun := adder(1, 2, 3, 2, 1, 2, 3, 4)
	fmt.Println(mysum, adderLen, fun)
}

func superman() {
	fmt.Println("I am Superman")
}

func multiple(v1 int, v2 int) int { //int outside the function arguments represents the output datatypes
	return v1 * v2
}

/*
Note -> No function overloading in go
*/
func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	length := len(values)
	mystr := "Just for fun"
	return sum, length, mystr
}
