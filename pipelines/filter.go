package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Msg string

// Listen simulates an infinite stream of messages, pushing them down an output
// channel.
func Listen(out chan<- Msg) {
	for {
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
		if rand.Intn(10) < 6 {
			out <- "foo"
		} else {
			out <- "bar"
		}
	}
}

// Filter just filters out "bar" messages.
func Filter(in, out chan Msg) {
	for {
		msg := <-in
		if msg == "bar" {
			continue // drop
		}
		out <- msg
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
func Store(in <-chan Msg) {
	for {
		msg := <-in
		fmt.Println(msg) // mock storage
	}
}

func main() {
    // seed the randomizer to randomize output between calls
	rand.Seed(time.Now().UnixNano())

	a := make(chan Msg)     // initial receiving channel
	b := make(chan Msg)     // -> a -> b -> c
	c := make(chan Msg)     // final outbound channel

	go Listen(a)            //   -> a
	go Filter(a, b)         // a -> b
	go Enrich(b, c)         // b -> c
	go Store(c)             // c ->

	time.Sleep(1 * time.Second)
}
