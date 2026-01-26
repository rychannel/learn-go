package main
import "fmt"

type car struct {
	color string
	body  string
}

func (c car) accelerate() string {
	return "accelerating--->"
}

func main() {
	porsche := car{
		color: "red",
		body: "coupe",
	}

	bently := car{
		color: "green",
		body: "saloon",
	}

	fmt.Println("Porsche paint is", porsche.color)
	fmt.Println("Porsche body type is", porsche.body)
	fmt.Println("Porsche is", porsche.accelerate())

	fmt.Println("Bently paint is", bently.color)
	fmt.Println("Bently body type is", bently.body)
	fmt.Println("Bently is", bently.accelerate())
}