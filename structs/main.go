package main

import "awesomeProject/structs/computer"
import "fmt"

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}
