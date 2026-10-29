package main 
import "fmt"

func main () {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]

	fmt.Println(arr)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 90, 100, 110)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}