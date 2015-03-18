package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    quit := make(chan bool) // HL
    c := boring("Joe", quit)
    for i := 10; i >= 0; i-- { fmt.Println(<-c) }
    quit <- true // HL
}

func boring(msg string, quit <-chan bool) <-chan string {
    c := make(chan string)
    go func() { // HL
        for i := 0; ; i++ {
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            select {
            case c <- fmt.Sprintf("%s: %d", msg, i):
                // do nothing
            case <-quit: // HL
                return
            }
        }
    }()
    return c
}
