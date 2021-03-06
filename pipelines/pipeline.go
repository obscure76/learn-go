package main

import "fmt"

// stage 1
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//stage 2
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//stage 3
func main() {
	//c := gen(2, 3)
	//out := sq(c)

	// Consume the output.
	//fmt.Println(<-out) // 4
	//fmt.Println(<-out) // 9

	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
