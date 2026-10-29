package main
import "fmt"

func main ()  {
	codes:= map[string]string{"en": "English", "fr" : "French", "hi" : "Hindi"}
    fmt.Println(codes)
	fmt.Println(len(codes)) 		//determine length of the map using len() function

	fmt.Println(codes["hi"])		//we can access the item of a map, by referring to it's key name
	fmt.Println(codes["en"])



}