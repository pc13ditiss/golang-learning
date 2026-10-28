package main

import "fmt"

func main (){
	var city string = "Bengaluru"
	{
		state:= "Karnataka"
		city := "Raigarh"
		fmt.Print(" The city ", city, " ", " is located in ", state, "\n")
	}
	fmt.Print(" The city ", city, " ",  "\n")
}

// the scope of the variable declared in the outer block can be accessed by the inner block
// but the print statement cannot access the declared variables in the inner block
//we say is local variables " declared inside" the function or block
// global variables "declared outside " the function or block