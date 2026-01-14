package main
import "fmt"

func main() {
	for i:=1; i<=3; i++ {
		for j := 1; j <=3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("\t\tBreaking out of inner loop when i=",i, "j=",j)
				break
			}
			if i==3 && j==2 {
				fmt.Println("\t\tContinuing inner loop when i=",i, "j=",j)
				continue
			}
			fmt.Println("Running i=",i, "j=",j)
		}
	}
}