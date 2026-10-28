//length of the array refers to the number of elements stored in the array
//we can find the length of an array using built-in length function len()
//this len() takes array as an input and returns us the size of the array

package main
import "fmt"

func main () {
	var fruits [2]string = [2]string{"apple", "banana"}    //normal string initialization
	fmt.Println(len(fruits))

	marks := [3]int {10, 20, 30}
	fmt.Println(len(marks))                                    // shorthand initialization

	name := [...]string{"Aaron", "Bob", "Chris"}         // ellipses initialization
	fmt.Println(name)
	fmt.Println(len(name)) 
}