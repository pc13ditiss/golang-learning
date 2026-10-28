package main 
import "fmt"

func main () {
	var grades [5]int = [5]int{70, 90, 80, 40, 50}
	for index, element :=range grades{                     //index and element are variables here
		fmt.Println(index, "=>", element)
	}
}