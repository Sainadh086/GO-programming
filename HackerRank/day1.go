package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var s2 string
    var num uint64
    var f float64

    // Read and save an integer, double, and String to your variables
    fmt.Scanf("%d\n%f\n", &num, &f)
    scanner.Scan()
    s2 = scanner.Text()

    //output
    fmt.Println(i + num)
    fmt.Printf("%.1f\n", d + f)
    fmt.Print(s + s2)
}
