package main

import "fmt"

func main () {
	var money int = 699
	var name string = "Sonyliv"
	var isCheck bool = true
	var amount float32 = 699.00

	fmt.Printf("var money = %v is of type %T \n" , money , money)
	fmt.Printf("var name = %v is of type %T \n" , name , name)
	fmt.Printf("var isCheck = %v is of type %T \n" , isCheck , isCheck)
	fmt.Printf("var amount = %v is of type %T \n" , amount , amount)
}

//value = literal (can also be said)