package main
import "fmt"

func main() {
	a := 8
	b := 4

	fmt.Println("Addition:\t", (a + b))      // 12
	fmt.Println("Subtraction:\t", (a - b))      // 4
	fmt.Println("Multiplication:\t", (a * b))   // 32
	fmt.Println("Division:\t", (a / b))         // 2
	fmt.Println("Modulus:\t", (a % b))          // 0
	a++
	fmt.Println("Increment:\t", a)          // 9
	b--
	fmt.Println("Decrement:\t", b)          // 3
}