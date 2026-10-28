package main 
import "fmt"

func main () {
	var grades [5]int = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(grades); i++{
		fmt.Println(grades[i])
	}
}