package main 
import "fmt"

func main () {
	arr := [10]int{85, 75, 70, 80, 90, 40, 60, 90, 40, 80}

	slice := arr[1:9]
	fmt.Println(slice)

	sub_slice := slice[0:3]
	fmt.Println(sub_slice)
}