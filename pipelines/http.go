package main

import (
	"fmt"
	"net/http" // production-grade HTTP server
)

type Msg struct {
	Data string    // 'data' parameter extracted from form values
	Done chan bool // signal channel to request handler
}

// Listen now constructs messages and passes them down the channel when it
// receives an incoming POST request with a `data` field in the form data.
// We're ppassing messages as pointers, now, because the stages can modify 
// the message.
func Listen(out chan *Msg) {
	h := func(w http.ResponseWriter, r *http.Request) {
		msg := &Msg{
			Data: r.FormValue("data"),
			Done: make(chan bool),
		}
		out <- msg

		success := <-msg.Done // wait for done signal
		if !success {
			w.Write([]byte(fmt.Sprintf("aborted: %s", msg.Data)))
			return
		}
		w.Write([]byte(fmt.Sprintf("OK: %s", msg.Data)))
	}

	http.HandleFunc("/incoming", h)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil) // blocks
}

func Filter(in, out chan *Msg) {
	for {
		msg := <-in
		if msg.Data == "bar" {
			msg.Done <- false
			continue
		}
		out <- msg
	}
}

func Enrich(in, out chan *Msg) {
	for {
		msg := <-in
		msg.Data = "☆ " + msg.Data + " ☆"
		out <- msg
	}
}

func Store(in chan *Msg) {
	for {
		msg := <-in
		fmt.Println(msg.Data)
		msg.Done <- true
	}
}

func main() {
	a := make(chan *Msg)     // initial receiving channel
	b := make(chan *Msg)     // -> a -> b -> c
	c := make(chan *Msg)     // final outbound channel

	go Listen(a)            //   -> a
	go Filter(a, b)         // a -> b
	go Enrich(b, c)         // b -> c
	go Store(c)             // c ->

	select {}               // block forever without spinning the CPU
}
