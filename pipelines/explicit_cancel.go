package main

//https://blog.golang.org/pipelines
import "sync"
import "fmt"
import "time"
import "math/rand"

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
			time.Sleep(time.Duration(rand.Intn(1000000)))
		}
		close(out)
	}()
	return out
}

func mergeV1(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		fmt.Println("Closing out channel in mergeV1")
		close(out)
	}()

	fmt.Println("Returning out from mergeV1")
	return out
}

/*
The value type of done is the empty struct because the value doesn't matter: it is the receive event that indicates the
send on out should be abandoned. The output goroutines continue looping on their inbound channel, c, so the upstream
stages are not blocked. (We'll discuss in a moment how to allow this loop to return early.)

Problem:
each downstream receiver needs to know the number of potentially blocked upstream senders and arrange to signal those
senders on early return. Keeping track of these counts is tedious and error-prone.
*/
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2)
	out := mergeV1(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}
}
