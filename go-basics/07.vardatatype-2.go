package main 

import (
"fmt"
"reflect"
)

func main () {
	fmt.Println("Type:  \n", reflect.TypeOf(699))
	fmt.Println("Type:  \n", reflect.TypeOf("SonyLiv"))
	fmt.Println("Type:  \n", reflect.TypeOf(true))
	fmt.Println("Type: \n", reflect.TypeOf(699.00))
}