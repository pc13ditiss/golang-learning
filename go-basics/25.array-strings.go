package main
import "fmt"

func main () {
	var fruits [2]string = [2]string{"apple", "banana"}    //normal string initialization
	fmt.Println(fruits)

	marks := [3]int {10, 20, 30}
	fmt.Println(marks)                                    // shorthand initialization

	name := [...]string{"Aaron", "Bob", "Chris"}         // ellipses initialization
	fmt.Println(name)
}