package main

import "fmt"

func main () {
	var name, surname string
	fmt.Println("enter you name")
	fmt.Scanf ("%s", &name)
    fmt.Println("enter you surname")
	fmt.Scanf ("%s", &surname)
	fmt.Println("Hey ", name, "", surname)
}

// %s = string specifier
// %t = boolean specifier
// %d = integer specifier
/* scanf function returns 2 values 
   count = the number of arguments that function counts to
   err = any error thrown during the execution of the function */