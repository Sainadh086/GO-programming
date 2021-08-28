// implementing slices
package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"a", "b", "c", "d", "e"}
	fmt.Println(things)

	things = append(things, "f")
	fmt.Println(things)

	things = append(things[1 : len(things)-1])

	myints := []int{4, 7, 2, 6, 8, 1}
	isSorted := sort.IntsAreSorted(myints)
	fmt.Println("Are ints Sorted", isSorted)
	sort.Ints(myints) //sorting integers
	fmt.Println("Sorted Values", myints)
}
