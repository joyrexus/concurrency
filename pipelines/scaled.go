package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Msg string

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

func Filter(in, out chan Msg) {
	for {
		msg := <-in
		if msg == "bar" {
			continue // drop
		}
		out <- msg
	}
}

func Enrich(in, out chan Msg) {
	for {
		msg := <-in
		msg = "☆ " + msg + " ☆"
		out <- msg
	}
}

func Store(in <-chan Msg) {
	for {
		msg := <-in
		fmt.Println(msg) // mock storage
	}
}

func main() {
    // Seeding the randomizer randomizes output between calls.
	rand.Seed(time.Now().UnixNano())

	a := make(chan Msg) // initial receiving channel
	b := make(chan Msg) // a -> b -> c
	c := make(chan Msg) // final outbound channel

    // Scaling the actors for a stage increases the concurrency of the program.
	go Listen(a)
	go Filter(a, b)
	go Enrich(b, c)
	go Enrich(b, c)
	go Enrich(b, c)
	go Store(c)

	time.Sleep(1 * time.Second)
}
