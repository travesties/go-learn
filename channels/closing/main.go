package main

import (
	"fmt"
)

func fibonacci(n int, ch chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}

	close(ch)
}

func main() {
	ch := make(chan int, 10)

	// Use channel capacity to cap the number of generated values.
	fibonacci(cap(ch), ch)

	// Since fibonacci is a sender/closer, we know that this channel
	// is closed and ready for reading. We can use range to read each
	// value sent to the channel, and the loop will terminate once
	// the last value is read.
	for val := range ch {
		fmt.Println(val)
	}
}
