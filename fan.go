package main

import (
    "fmt"
    "time"
)

func say(msg string) <-chan string { // Returns receive-only channel of strings.
    c := make(chan string)
    go func() { // We launch the goroutine from inside the function.
        for {
            c <- msg
            time.Sleep(1 * time.Second)
        }
    }()
    return c // Return the channel to the caller.
}

// multiplex two channels
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { 
        for { 
            select {
            case s := <-input1: c <- s
            case s := <-input2: c <- s
            }
        }
    }()
    return c
}

func main() {
    hi := say("hi!") // Function returning a channel.
    ho := say("ho!") // Function returning a channel.
    c := fanIn(hi, ho)
    for i := 0; i < 5; i++ {
        fmt.Println(<-c)
    }
    fmt.Println("You're boring; I'm leaving.")
}
