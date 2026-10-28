package main
import "fmt"

func main () {
	var x, y int = 10, 20
	x += y                          x = x+y here
	fmt.Println(x)
}

// right side where the value is written(here y) and left side where you want the result (here x)