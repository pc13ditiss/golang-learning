package main
import "fmt"

func main () {
	arr := [10]int{10, 15, 20, 25, 30, 35, 40, 45, 50}

	slice := arr[1:4]
	slice_2 := arr[4:8]
	fmt.Println(slice)
	fmt.Println(slice_2)

	slice_3 := append(slice, slice_2...)
	fmt.Println(slice_3)
}