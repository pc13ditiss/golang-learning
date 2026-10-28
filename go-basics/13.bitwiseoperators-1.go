/* bitwise AND &     ------> Returns 1 only if **both bits are 1
   bitwise OR |      ------> Returns 1 if **at least one bit is 1
   bitwise XOR ^.    ------> Returns 1 if **the bits are different
   right shift >>.   ------> **Halves the number by shifting bits right
   left shift <<.    ------> **Doubles the number by shifting bits left

*/
package main
import "fmt"

func main () {
	var x, y int = 12,25
	z := x & y
    fmt.Println(z)
}
/*
   0000 1100 
 & 0001 1001
   ----------
   0000 1000
*/