package main
import "fmt"

func main(){
	//Print statement
	fmt.Println("Defining the variables")
	// writing  variables in different formats
	var name string = "Sainadh"
	fmt.Println(name) //fmt has many other operations and functions
	var friend string
	friend = "Golang"
	fmt.Println(friend)

	// this special case and we no need to give datatype
	sister := "Kavya"
	fmt.Println(sister)

	var age int = 22
	fmt.Println(age)

	year := 1999
	fmt.Println(year)

	//printing the value and datatype
	fmt.Printf("%v, %T",age,age)
}
