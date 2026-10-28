/* strconv package
Itoa() 
converts integer to string, returns one value - string formed with the given integer. 

*/

package main 

import (
	"fmt"
	"strconv"
	"reflect"
)

func main () {
	var i int = 42
	var s string = strconv.Itoa(i) //convert int to string 
	fmt.Printf("%q \n", s)

	fmt.Println("Type: ", reflect.TypeOf(s))
}

/* 

It prints a string or character wrapped in quotes ("" or '') and 
automatically handles and escapes any special characters (like tabs, newlines, or quotes inside the string).

It's very handy for debugging because it makes hidden characters visible.
 For example, if a string contains a newline (\n), %q will print the literal 
 text \n instead of actually breaking the line, so you can see it clearly:*/
