package main
import "fmt"

func main () {
	var fruits [5]string = [5]string{"apple", "banana", "grapes", "cherry", "oranges"}    //normal string initialization
	fmt.Println(fruits)

	fruits[1] = "mango"
	fmt.Println(fruits)

}