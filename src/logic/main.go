package main
import "fmt"

func main() {
	var yes, no bool = true, false

	result := ( yes && no )
	fmt.Println("Logical AND:\t\tyes && no\t", result)	  // false

	result = ( yes || no )
	fmt.Println("Logical OR:\t\tyes || no\t", result)	  // true

	result = !yes
	fmt.Println("Logical NOT:\t\t!yes\t\t", result)	  // false
}