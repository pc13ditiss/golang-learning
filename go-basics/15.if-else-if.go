package main

import "fmt"

func main () {
	var fruit string = "grapes"
	if fruit == "apple" {
		fmt.Println("The fruit is an apple")
	}	else if fruit == "orange" {
		fmt.Println("The fruit is not an apple")
	}   else if fruit == "grapes" {
		fmt.Println("I love grapes")
	}   else {
		fmt.Println("I don't have an appetiite")
	}
}