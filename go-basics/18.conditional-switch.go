package main
import "fmt"

func main () {
	var a, b int = 10, 30                           //here we do not give break keyword, it breaks automatically
	switch {
	case a+b == 50:
		fmt.Println("Then value is 50")
	case a+b <= 50:
		fmt.Println("Then value is less than 50")
	default:
		fmt.Println("Then value is ou of context")
	}
}