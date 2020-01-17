package main

import (
	"fmt"
	"time"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
	fmt.Println("squareop", sum)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	time.Sleep(3000*time.Millisecond)
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
	fmt.Println("cubeop", sum)
}


func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)



	number = 589
	sqrch = make(chan int)
	cubech = make(chan int)
	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes = <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}