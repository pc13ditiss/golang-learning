
package main
import "fmt"

func main () {
   var i int = 90
   var co float64 = 699.01
   var f float64 = float64(i)
   var c int = int (co)

   fmt.Printf("%.2f \n", f)
   fmt.Printf("%v \n", c)

}

//Typecasting = the process of converting one data type to another is called typecasting.

// The data types can be converted to another data types but this does not gurantee the value of the variable remains intact.
/*
Here is the short difference between fmt.Println and fmt.Printf in Go:

fmt.Println

What it does: Prints values as-is, adds spaces between arguments automatically, and automatically adds a newline at the end.

Format verbs: Does not support format verbs (like %d or %.2f).

fmt.Printf

What it does: Prints text using format verbs (like %s, %d, %.2f) to control how data looks.

Newline: Does not automatically add a newline—you must add \n manually at the end of the string.
*/