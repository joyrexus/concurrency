package main

import (
	"fmt"
	"net/http" // production-grade HTTP server
)

// Msg type now contains various message features, including the path on which
// the http POST request was received, the `data` parameter from the form
// values in the POST, and a signal channel for notifying the handler that our
// pipeline is finished.  That is, we need to signal the HTTP handler to write 
// a response to the requesting client and close the connection (see `Listen`
// function).  
//
// The signaling channel also indicates the message status: it receives a 
// value of `false` when a message gets filtered and `true` when stored.
type Msg struct {
    Path string    // route receiving the data
	Data string    // 'data' parameter extracted from form values
	Done chan bool // signal channel to request handler
}

// Listen now constructs messages and passes them down the channel when it
// receives an incoming POST request with a `data` field in the form data.
// We're passing messages as pointers, now, because the stages can modify 
// the message.
func Listen(out chan *Msg) {
	h := func(w http.ResponseWriter, r *http.Request) {
		msg := &Msg{
            Path: r.RequestURI,
			Data: r.FormValue("data"),
			Done: make(chan bool),
		}
		out <- msg

		success := <-msg.Done   // wait for done signal
		if !success {           // message was filtered
            err := fmt.Sprintf("aborted: %s received on %s", msg.Data, msg.Path)
			w.Write([]byte(err))
			return
		}
		w.Write([]byte(fmt.Sprintf("OK: %s", msg.Data)))
	}

	http.HandleFunc("/kosher", h)       // accept messages from this route
	http.HandleFunc("/suspect", h)      // filter messages from this route
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil) // blocks
}

func Filter(in, out chan *Msg) {
	for {
		msg := <-in
		if msg.Path == "/suspect" {
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
