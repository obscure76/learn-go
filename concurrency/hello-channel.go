package main

import (
	"fmt"
	"time"
)

func helloLocal(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(8000*time.Millisecond)
	fmt.Println("Hello world goroutine after wait")
	done <- true
}
func main() {
	done := make(chan bool)
	go helloLocal(done)
	<-done
	fmt.Println("main function")
}