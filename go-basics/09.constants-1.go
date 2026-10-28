package main

import "fmt"

func main () {
	const name = "vipul"
	const is_man = true
	const age = 28

	fmt.Printf("%v: %T\n", name, name)
	fmt.Printf("%v: %T\n", is_man, is_man)
	fmt.Printf("%v: %T\n", age, age)
}

/*we cannot assign or reassing the variable in constants
const name = "vipul"
name = "xaylar" 

&&

we need to assign value at the time of declaration not later, like
const name
name = " vipul"

both of these conditions will throw an error*/

// const name := "vipul" --> this is a shorthand variable assignment operator


//untyped constants is above - it is flexible and flexibility is lost in untyped constants
// typed constants --> const name string = "vipul"