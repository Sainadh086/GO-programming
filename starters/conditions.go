// Basic Implementation of conditions with go
// Scenario taking rating as input and printing a messsage based on the rating

package main

import (
	"bufio"   // for reading from stdin
	"fmt"     // we use fmt for printing
	"os"      // for standard input reading
	"strconv" // for converting string to numerical values
	"strings" // for splitting string
	"time"    // time object
)

// we use time object also
func main() {
	// Frontend -> Input
	var name, rating string
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')

	fmt.Println("Please rate our Dosa center between 1 to 5")
	reader = bufio.NewReader(os.Stdin)
	rating, _ = reader.ReadString('\n')
	ratingNum, _ := strconv.Atoi(strings.TrimSpace(rating))

	// Backend -> Output
	fmt.Printf("Hello %s, \n Thanks for rating our dosa center with %v,\n Your rating was recorded in our system at %v\n",
		name, ratingNum, time.Now().Format(time.Stamp))

	//implementing conditions
	if ratingNum == 5 {
		fmt.Println("Bonus for team for 5 star services")
	} else if ratingNum == 4 || ratingNum == 3 { // using or condiotion, and -> &&
		fmt.Println("We are always improving")
	} else if ratingNum < 3 {
		fmt.Println("We require serious work improvement from our side")
	}
}

// Basic Implementation of conditions with go
/*
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating between 1 to 5 : ")
	rating_input, _ := reader.ReadString('\n')
	rating_input = strings.TrimSpace(rating_input)
	rating, _ := strconv.ParseFloat(rating_input, 64)
	if rating >= 0 && rating <= 5 {
		if rating == 0 {
			fmt.Println("You are a terrible person")
		} else if rating == 1 {
			fmt.Println("You are a bad person")
		} else if rating == 2 {
			fmt.Println("You are a mediocre person")
		} else if rating == 3 {
			fmt.Println("You are a good person")
		} else if rating == 4 {
			fmt.Println("You are a great person")
		} else if rating == 5 {
			fmt.Println("You are a fantastic person")
		}
	} else {
		fmt.Println("Invalid rating")
	}
}*/
