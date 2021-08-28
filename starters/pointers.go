//  using pointers
package main

import "fmt"

func main() {
	// General definition of pointers
	//var p *int

	var age int = 40
	p := &age // p is a pointer to the memory address where age is stored

	fmt.Println(age)
	fmt.Println(p)
	fmt.Println(*p) // read the value at the memory address p is pointing to

}
