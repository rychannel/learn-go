package main

import (
	"cube"
	"fmt"
)

func main() {
	var box cube.Dims

	box.SetSize(2, 4, 6)

	fmt.Printf("Volume:%v \n", box.GetVolume())
	fmt.Printf("Area:%v \n", box.GetArea())

	fmt.Println("Width:", box.width)
	fmt.Println("Area:", box.getArea())
}
