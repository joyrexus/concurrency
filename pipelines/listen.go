package main

import (
    "fmt"
    "time"
    "math/rand"
)

type Msg string

// Seed seeds the random function, which allows us
// to randomize the output between calls.
func Seed() {
    rand.Seed(time.Now().UnixNano())
}

// Listen simulates an infinite stream of messages, pushing
// them down an output channel.
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

// Print takes an input channel and prints values as they
// come down the channel.
func Print(in <-chan Msg) {
    for {
        msg := <-in
        fmt.Println(msg)
    }
}

func main() {
    Seed()
    c := make(chan Msg) 
    go Listen(c)
    go Print(c)
    time.Sleep(1 * time.Second)
}

