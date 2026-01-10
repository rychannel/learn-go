package main
import "fmt"

func main() {
	var zero, num, max int = 0, 0, 1
	var up, dn byte = 'A', 'a'

	result := ( num == zero)
	fmt.Println("Equality:\t\tnum ==zero\t", result)		  // true

	result = ( up == dn )
	fmt.Println ("Inequality:\t\tup == dn\t", result)		  // false

	result = (max != zero)
	fmt.Println("Inequality:\t\tmax != zero\t", result)	  // true

	result = ( zero > max )
	fmt.Println("Greater than:\t\tzero > max\t", result)	  // false

	result = ( max <= zero )
	fmt.Println("Less than or equal:\tmax <= zero\t", result) // false
}