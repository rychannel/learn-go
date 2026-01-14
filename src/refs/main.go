package main
import "fmt"

func square( num *int ) {
	fmt.Println("\t\tRecieved Address", num)

	*num = *num * *num
	fmt.Println("\t\tModified Original:", *num )
}

func main() {
	num := 5
	fmt.Println("Original:", num)
	square( &num )
	fmt.Println("Original:", num)
}