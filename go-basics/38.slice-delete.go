package main
import "fmt"

func main () {
	arr := [10]int{10, 15, 20, 25, 30, 35, 40, 45, 50}

	i := 3                             // actual number or order
	fmt.Println(arr)
	slice := arr[:i]
	slice_2 := arr[i+1:]              //first element (i) is index starts from 0, second number is the actual number or order starts from 1

	fmt.Println(slice)
	fmt.Println(slice_2)

	new_slice := append(slice, slice_2...)
	fmt.Println(new_slice)                  //Here number 25 in index 3 is deleted
}