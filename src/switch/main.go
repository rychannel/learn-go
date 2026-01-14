package main
import "fmt"

func main() {
	num := 2
	char := 'B'

	switch num {
		case 1:
			fmt.Println("\nNumber is 1")
		case 2:
			fmt.Println("\nNumber is 2")
		case 3:
			fmt.Println("\nNumber is 3")
		default:
			fmt.Println("\nNumber is not 1, 2 or 3")
	}

	switch char {
		case 'A':
			fmt.Println("\nCharacter is A")
		case 'B':
			fmt.Println("\nCharacter is B")
		case 'C':
			fmt.Println("\nCharacter is C")
		default:
			fmt.Println("\nCharacter is not A, B or C")
	}
}