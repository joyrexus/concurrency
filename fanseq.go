package main

import (
    "fmt"
    "time"
)

// Multiplex the input channels.
func fanIn(inputs ... <-chan Message) <-chan Message {
    c := make(chan Message)
    for i := range inputs { 
        input := inputs[i]
        go func() {         // Launch goroutine for each input.
            for {
                c <- <-input 
            }
        }()
    }
    return c
}
type Message struct {
    str string
    resume chan bool
}

// Returns receive-only channel of Messages.
func say(msg string) <-chan Message { 
    c := make(chan Message)
    wait := make(chan bool)
    go func() {
        for {
            c <- Message{msg, wait}
            time.Sleep(1 * time.Second)
            <-wait  // block until receiving value on the channel
        }
    }()
    return c // Return the channel to the caller.
}

func main() {
    hi := say("hi!")    // Function returning a channel.
    ho := say("ho!")    // Function returning a channel.
    c := fanIn(hi, ho)  // Multiplex both channels onto one channel.
    for i := 0; i < 5; i++ {
        msg1 := <-c; fmt.Println(msg1.str)
        msg2 := <-c; fmt.Println(msg2.str)
        msg1.resume <- true
        msg2.resume <- true
    }
    fmt.Println("You're boring; I'm leaving.")
}
