package main
import "fmt"

func main () {
	var x,y int = 10, 5
	x *=y                       // x = x*y here
	fmt.Println(x)
}