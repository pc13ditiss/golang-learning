package main
import "fmt"

func main () {
	var x, y int = 12,25
	z := x ^ y                 //result of XOR is 1 if two bits are opposite, 0 if both bits are same
    fmt.Println(z)
}

/*
   0000 1100 
 ^ 0001 1001
   ----------
   0001 0101
*/