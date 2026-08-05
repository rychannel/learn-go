package main

import "fmt"

func countDn(num int) {
	if num < 1 {
		fmt.Println("\t\t\t\tLift Off!")
	} else {
		fmt.Println("\t\t\tCountdown", num)
		num--
		countDn(num)
	}
}

func main() {
	countDn(10)
}
