package main

import "fmt"

func main() {
	area := func(length, width int) int {
		return length * width
	}

	fmt.Printf("area Type: %T \n", area)
	fmt.Println("area 1:", area(10, 4))
	fmt.Println("area 2:", area(12, 5))

	counter := func() func() int {
		num := 0
		return func() int {
			num++
			return num
		}
	}()
	fmt.Printf("counter Type: %T \n", counter)
	fmt.Println("count 1:", counter())
	fmt.Println("count 2:", counter())
	fmt.Println("count 3:", counter())
}
