// Atoi() converts string to integer 
// returns two values - corresponding integer, error (if any)

package main

import (
	"fmt"
	"strconv"
)

func main () {
	var s string = "vipul"
	var t string = "699"

	i, err := strconv.Atoi(s)
	p, err := strconv.Atoi(t)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

	fmt.Printf("%v, %T \n", p, p)
	fmt.Printf("%v, %T \n", err, err)

}

//In Go, %T is a format verb used with fmt.Printf to print the data type of a variable.

//It is a clean, built-in shortcut for what you were trying to do with the reflect package