First stab at creating examples of basic concurrency patterns.

Sources for all examples listed below.

---

For a quick review of goroutines and channels, see the concurrency sections of
the golang [tour](http://tour.golang.org/concurrency/1), [bootcamp](http://www.golangbootcamp.com/book/concurrency), or [intro-book](http://www.golang-book.com/10/index.htm).

*Channels are concurrent-safe queues that are used to safely pass messages
between Go’s lightweight processes (goroutines). The message-passing style they encourage permits the programmer to safely coordinate multiple concurrent tasks with easy-to-reason-about semantics and control flow that often trumps the use of callbacks or shared memory.* — Alan Shreve


## Articles

* [Timing Out, Moving On](http://blog.golang.org/go-concurrency-patterns-timing-out-and)
* [Pipelines and Cancellation](http://blog.golang.org/pipelines)
* [Context](http://blog.golang.org/context)
* [API Patterns](https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/) - principles of designing Go APIs with channels
* [Principles of designing Go APIs with channels](https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/)


## Slides

* [Concurrency is not Parallelism](http://talks.golang.org/2012/waza.slide#1)
* [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide)
  * [example source files](https://github.com/golang/talks/tree/master/2012/concurrency/support)
* [Advanced Go Concurrency Patterns](https://talks.golang.org/2013/advconc.slide)
* [Cancelation, Context, and Plumbing](https://talks.golang.org/2014/gotham-context.slide)
* [Querying multiple backends](http://peter.bourgon.org/go-do/#16)
* [Pipelined data processing](http://peter.bourgon.org/go-do/#24)
