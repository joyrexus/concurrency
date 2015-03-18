package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Msg string

// Listen simulates an infinite stream of messages, pushing them down an output
// channel.
func Listen(out chan Msg) {
	for {
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
		if rand.Intn(10) < 6 {
			out <- "foo"
		} else {
			out <- "bar"
		}
	}
}

// Enrich reads a single message from the input channel, processes it, 
// and pushes the result down the output channel.
func Enrich(in, out chan Msg) {
	for {
		msg := <-in
		msg = "☆ " + msg + " ☆"
		out <- msg
	}
}

// Store simulates writing the message somewhere.
func Store(in chan Msg) {
	for {
		msg := <-in
		fmt.Println(msg) // store to stdout
	}
}

func main() {
    // seed the randomizer to randomize output between calls
	rand.Seed(time.Now().UnixNano())

	// channels for our pipeline (msg -> a -> b -> stdout)
	a := make(chan Msg)     // these are unbuffered channels
	b := make(chan Msg)     // with automatic backpressure

    // wire the stages together, and launch each stage as a goroutine
	go Listen(a)
	go Enrich(a, b)
	go Store(b)

	time.Sleep(2 * time.Second)
}
