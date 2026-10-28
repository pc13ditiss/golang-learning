package main
import "fmt"

func main () {
	var f int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &f)
	for i := 1; i<11; i++{
        fmt.Println(i*f)
	}
}