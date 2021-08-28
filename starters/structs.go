// implementing structures
// It helps us to define a new data structure

// Note: go doesn't support inheritance
package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	// one way easy to use
	sai := User{"Sai", "sai@go.com", 22}
	fmt.Printf("%+v\n", sai)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@go.com"
	sam.age = 22
	fmt.Printf("%+v\n", sam)
	// as we used new to create the user DS, it only refer to the data

}
