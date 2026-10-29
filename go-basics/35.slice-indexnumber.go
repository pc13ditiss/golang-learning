package main
import "fmt"

func main () {
	arr := [5]int{10, 30, 40, 60, 50}

	slice := arr[ : 3]
	fmt.Println(arr)
	fmt.Println(slice)

	slice[1] = 200
	
	fmt.Println("array after mofification = ", arr)
	fmt.Println("slice after mofification = ", slice)

}