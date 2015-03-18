package main

import "fmt"

func main() {
    c := make(chan int)
    go upto(10, c)
    consume(c)
}

func upto(n int, c chan int) {
    for i := 0; i < n; i++ {
        c <- i // put data onto channel
    }
    close(c)
}

func consume(c chan int) {
    for i := range c {
        fmt.Println(i)
    }
}
