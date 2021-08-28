package main

import (
	"fmt"
)

func main() {
	// using map data structure which is key,value pair
	// new - only allocates - not initializes the memory
	// make - allocates and initializes the memory

	score := make(map[string]int)
	score["sai"] = 89
	fmt.Println(score)

	score["varshan"] = 34
	score["yogesh"] = 99
	fmt.Println(score)
	getVarshanScore := score["varshan"]
	fmt.Println(getVarshanScore)

	//delete
	delete(score, "varshan")
	fmt.Println(score)

	//using for loop
	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
