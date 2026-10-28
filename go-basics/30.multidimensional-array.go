//multidimensional array is called arrays of array

package main
import "fmt"

func main () {
	arr := [3][2]int{{2,4}, {6,8}, {11,9}}  //here 3 refers size of the array, 2 refers size of each element itself
	fmt.Println(arr)
	fmt.Println(arr[2][0])
}