package main

import (
    "fmt"
    "time"
    "math/rand"
)

// A simple generator yielding an integer starting 
// with n, incremented after each yield.
func counter(n int) <-chan int {
    c := make(chan int)
    go func() { 
        for {
            c <- n; n++
            time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
        }
    }()
    return c
}

func main() {
    c := counter(4)
    loop:
        for i := 0; i < 10; i++ {
            select { 
            case v := <-c: 
                fmt.Println(v)
            // timeout if counter channel doesn't return a value within a second
            case <-time.After(1 * time.Second): 
                fmt.Println("You're too slow!")
                break loop
            }
        }

    fmt.Println("Continue pulling values for 4 seconds ...")
    timeout := time.After(4 * time.Second)  // timeout after 4 seconds
    for {
        select { 
        case v := <-c: 
            fmt.Println(v)
        case <-timeout:
            fmt.Println("4 seconds is up!")
            return
        }
    }
}
