package main
import "fmt"

func main () {
	var grades [3]int = [3]int{10, 20, 30} 
//	grades := [3]int{10, 20, 30}               //shorthand declaration of array
//	grades := [...]int{10, 20, 30}             //ellipses declaration of array, here we do notneed to define the length 
    fmt.Println(grades)
}