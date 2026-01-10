package main
import "fmt"

func main() {
	const pi = 3.14159

	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)

	fmt.Printf("Red: %v points \n", red)
	fmt.Printf("Blue: %v points \n", blue)
	fmt.Printf("Black: %v points \n", black)
}