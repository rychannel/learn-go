package main
import "fmt"

func main() {
	var a, b int

	a, b = 8, 4

	fmt.Println("Assigned Values:\ta =", a, ", b =", b)

	a += b
	fmt.Println("Add and Assign:\t\ta =", a)

	a -= b
	fmt.Println("Subtract and Assign:\ta =", a)

	a *= b
	fmt.Println("Multiply and Assign:\ta =", a)

	a /= b
	fmt.Println("Divide and Assign:\ta =", a)

	a %= b
	fmt.Println("Modulus and Assign:\ta =", a)
}