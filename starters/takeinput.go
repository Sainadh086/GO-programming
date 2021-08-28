//input and ouput operations

package main
import (
	"fmt"
	"bufio" //for multiple elements as input
	"os"
	"strconv"
	"strings"
)

func main(){

	/*var name string
	//reading from the cli only for single word
	fmt.Scanln(&name)
	fmt.Println("Your name is : ",name)
	*/
	/*var name string = "Sai"
	var a, _ = fmt.Println(name)
	fmt.Println(name)
	*/
	
	// reading from the cli for multiple words
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")
	//here we are using '_' to ignore the error
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is : ",name)
	

	//taking numerical input
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")
	rating, _ := reader.ReadString('\n')
	ratingNum, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)  //64 indicates that the number is of type float64, number of bits
	fmt.Println("Your rating is : ",ratingNum+2)
}
