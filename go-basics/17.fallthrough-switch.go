// fallthrough keyword is used in switch-case to force the execution flow to fall through the successive case block

package main
import "fmt"

func main () {
	var i int = 10
	switch i {
    	case -5: 
			fmt.Println(" i is -5")
		case 0: 
			fmt.Println(" i is 0")
		
		case 10: 
			fmt.Println(" i is 10")
			fallthrough.              // here fallthrough prints 10, 20 and default also 
		
		case 20: 
			fmt.Println(" i is 20")
			fallthrough
		
		default:
			fmt.Println(" i is neither of above ")
	}
}