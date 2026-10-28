package main
import "fmt"

const PI float64 = 3.14

func main () {
	var radius float64 
	fmt.Printf("Enter the radius of the circle \n")
	fmt.Scanf("%f", &radius)

	var area float64 = PI * radius * radius

	fmt.Printf("The area of the circle is = %.2f\n ", area)

}
