package main
import "fmt"

func main () {
	var x int = 212
	z := x >> 1             //result of right shift is half, the new replace number is added by 0 in vacant place
    fmt.Println(z)
}

/*
    1101 0100 = 212
 >> 01101 010 = 106
*/