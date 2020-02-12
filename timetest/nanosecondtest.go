package main

import "time"
import "fmt"

func main() {
	var timens time.Duration = 1000000
	fmt.Println(timens)
	var div = timens / time.Millisecond
	fmt.Println(div)
	fmt.Println(time.Second / time.Millisecond)
	var val = fmt.Sprint(time.Nanosecond / time.Nanosecond)
	fmt.Println(val)
}
