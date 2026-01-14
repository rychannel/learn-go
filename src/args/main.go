package main
import "fmt"

func square( num int ) {
	fmt.Println("\t\tRecieved Copy", num)

    num = num * num
    fmt.Println("\t\tModified Copy", num )
    
}

func main() {
	num := 5
	fmt.Println("Original", num)
	square( num )
	fmt.Println("Original", num)
}