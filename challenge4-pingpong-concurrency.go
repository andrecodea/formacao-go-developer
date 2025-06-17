package main

import (
	"fmt"
	"sync"
	"time"
)

// Produces messages and sends them through channels
func pingpong(c1, c2 chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Makes sure that wg.Done() is called after the function is executed
	// Indicates to WaitGroup that this goroutine is finished

	for i := 0; i < 10; i++ {
		c1 <- "ping" // Sends "ping" to channel c1 (blocking operation)
		c2 <- "pong" // Sends "pong" to channel c2 (blocking operation)

		// ⚠️ IMPORTANT: As the channels are unbuffered, each sending operation blocks until
		// someone receives the message.
	}

	// Closes the channels to indicate that there's no more data
	close(c1) // Indicates that c1 is not receiving anymore messages
	close(c2) // Indicates that c2 is not receiving anymore messages

	// When a channel is closed, receiving operations return zero and false at the second parameter
}

// Consumes channels's messages and prints them
func print_msg(c1, c2 chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Indicates that this goroutine is done, but only after the function executes

	// Infinite loop - only for when the channels are closed
	for {
		// Receives c1 with closing verification
		msg1, ok1 := <-c1 // msg1 = message received, ok1 = true if channel is open
		if !ok1 {         // If channel c1 is closed
			break // Break loop
		}
		fmt.Println(msg1)           // Prints "ping"
		time.Sleep(time.Second * 1) // Sleeps for 1 second

		// Receives c2 with closing verification
		msg2, ok2 := <-c2 // msg2 = message received, ok2 = true if channel is open
		if !ok2 {         // If channel c2 is closed
			break // Break loop
		}
		fmt.Println(msg2)           // Prints "pong"
		time.Sleep(1 * time.Second) // Sleeps for 1 second

		// 2 second total per cycle (ping + pong)
	}
}

func main() {
	// Creates two unbuffered channels for strings
	c1 := make(chan string) // Channel for "ping" messages
	c2 := make(chan string) // Channel for "pong" messages

	// Creates a WaitGroup to sync goroutines
	var wg sync.WaitGroup
	wg.Add(2) // Adds 2 goroutines to counter

	// Initializes producing goroutine in parallel
	go pingpong(c1, c2, &wg) // Passes WaitGroup pointer

	// Initializes consuming goroutine in parallel
	go print_msg(c1, c2, &wg) // Passes WaitGroup pointer

	// Blocks main goroutine until both goroutines call wg.Done(), which is defered
	wg.Wait()

	// Only executes after both goroutines are done
	fmt.Println("End!")
}
