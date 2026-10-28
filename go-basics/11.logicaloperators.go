package main
import "fmt"

func main () {
	var x int = 10
	var y int = 50

	fmt.Println((x<100) && (x > 5))
	fmt.Println((x<300) && (x > 12))

	fmt.Println((x<300) ||  (x > 12))
	fmt.Println((x<0) || (x > 12))

	fmt.Println(!(x>y))
	fmt.Println(!(true))
	fmt.Println(!(false))


}