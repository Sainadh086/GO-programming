package main

import "fmt"

func main() {
	// declaring arrays
	var numbers [3]string
	numbers[0] = "one"
	numbers[1] = "two"
	numbers[2] = "three"

	fmt.Println(numbers)

	// declaring and initializing arrays
	var colors = [3]string{"red", "green", "blue"}
	fmt.Println(colors[0])
	fmt.Println(len(colors))
}
