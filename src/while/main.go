package main
import "fmt"

func main() {

	counter :=1
	for counter <=5 {
		fmt.Println("While Loop Iteration", counter)
		counter++
	}

	i := 10
	for {
		fmt.Println("\t\t\tCountdown", i)
		i--
		if i == 0 {
			fmt.Println("\t\t\t\t\tLiftoff!")
			break
		}
	}
}