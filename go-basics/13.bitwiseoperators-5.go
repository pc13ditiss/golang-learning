package main
import "fmt"

func main () {
	var x int = 212
	z := x << 1             //result of left shift is double, the new replace number is added by 0 in vacant place only 1 bit here
    fmt.Println(z)
}

/*
    1101 0100 = 212
 << 101 01000 = 424
*/