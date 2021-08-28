// Even though GO supports OOPs concept then don't have classes, We can use structutre instead
package main

import "fmt"

//definig structure
type Book struct {
	Author   string
	BookName string
	Year     int
}

// method
func (book Book) print_details() {
	fmt.Println("The book " + book.BookName + " was published by " + book.Author)
	fmt.Println("It is published in the year ", book.Year)

}

func main() {
	b1 := Book{"Sai", "The power of positive thinking", 2021}
	b1.print_details()

	//entering multiple values
	list := []Book{
		Book{"yogesh", "Go Programming", 2020},
		Book{"varshan", "Everything about GO", 2021},
		Book{"balaji", "Cooking is an art", 2020},
	}

	for i := range list {
		list[i].print_details()
	}
}
