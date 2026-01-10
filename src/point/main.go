package main
import "fmt"

func main() {
	var num int = 20
	var ptr *int = &num

	fmt.Printf("num valud: %v type: %T \n", num, num)
	fmt.Printf("num address: %v type: %T \n\n", ptr, ptr)

	fmt.Printf("num via pointer: %v type: %T \n", *ptr, *ptr)
	fmt.Printf("pointer address: %v type: %T \n\n", &ptr, &ptr)

	*ptr = 100
	fmt.Printf("new num valude: %v type: %T \n", num, num)
}