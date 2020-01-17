package main

import "fmt"

func describeLocal(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describeLocal(s)
	i := 55
	describeLocal(i)
	strt := struct {
		name string
	}{
		name: "Midhun Chinta",
	}
	describeLocal(strt)
}
