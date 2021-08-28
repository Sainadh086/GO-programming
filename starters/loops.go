package main

import "fmt"

func main() {
	start := 1
	things := []string{"ball", "bat", "watch", "mobile"}

	//for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	// range for simpler use
	for i := range things {
		fmt.Println(things[i])
	}

	// we don't have while here
	for start < 5 {
		if start == 4 {
			//break //we cam also use continue
			goto mylabel
		}
		fmt.Println("Start is now : ", start)
		start += start
	}

	// using labels
mylabel:
	fmt.Println("Learning Loops in GO")
}
