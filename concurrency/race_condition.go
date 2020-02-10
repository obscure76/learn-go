package main

import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup, x *int) {
	*x = *x + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var x int
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &x)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
