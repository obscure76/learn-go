package main

import (
	"fmt"
	"sync"
)

func incrementLocal(wg *sync.WaitGroup, ch chan bool, xl *int) {
	ch <- true
	*xl = *xl + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	var xl int
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementLocal(&w, ch, &xl)
	}
	w.Wait()
	fmt.Println("final value of x", xl)
}
