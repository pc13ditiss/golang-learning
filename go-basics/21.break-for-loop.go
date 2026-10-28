package main
import "fmt"

func main () {
	for i := 1; i<11; i++{
		if i == 6 {
			break               // it print only 5th but skips the 6th
		}
        fmt.Println(i)
	}

}