package main

import (
    "fmt"
    "math/rand"
    "time"
)

func Init() {
    rand.Seed(time.Now().UnixNano())
}

type Server interface {
    Query(q string) string
}

type Alpha struct{}

func (a Alpha) Query(q string) string {
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    return fmt.Sprintf("alpha/%s", q)
}

type Beta string

func (b Beta) Query(q string) string {
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    return fmt.Sprintf("%s/%s", b, q)
}

type Replicas []Server

func (r Replicas) Query(q string) string {
    c := make(chan string, len(r))
    for _, server := range r {
        go func(s Server) { c <- s.Query(q) }(server)
    }
    return <-c
}

func QueryAll(q string, servers ...Server) []string {

    results := []string{}
    c := make(chan string, len(servers)) // buffered chan

    // query
    for _, server := range servers {
        go func(s Server) { c <- s.Query(q) }(server)
    }

    // aggregate
    for i := 0; i < cap(c); i++ {
        results = append(results, <-c)
    }
    return results
}

func main() {
    Init()

    r1 := Replicas{
        Beta("server-1"),
        Beta("server-2"),
        Beta("server-3"),
    }
    r2 := Replicas{
        Beta("server-4"),
        Beta("server-5"),
    }
    r3 := Replicas{
        Alpha{},
        Alpha{},
        Alpha{},
    }

    start := time.Now()
    results := QueryAll("foo", r1, r2, r3)
    fmt.Println(results)
    fmt.Println(time.Since(start))
}
