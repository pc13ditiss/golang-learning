package main
import "fmt"

func main () {
	for i := 1; i<11; i++{
		if i == 6 {
			continue               // it prints all values but skips the 6th
		}
        fmt.Println(i)
	}

}