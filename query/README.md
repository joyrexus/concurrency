Example of collecting data from multiple backends from Peter Bourgon's talk [Go Do](http://peter.bourgon.org/go-do/#16). See also [sample code](https://github.com/peterbourgon/peter-bourgon-org/tree/master/src/go-do) from the slides.

This example was apparently adapted from Sameer Ajmani's [Google Tech
Talk](http://youtu.be/4iAiS-qv26Q).

For context, we're imagining a set of backend servers (possibly replicas) that
can perform queries.  We want to broadcast a query to a set of backends, and
aggregate the responses.  The goal is to fire off the queries concurrently
(instead of synchronously) and aggregate the results in whatever order they
happen to come in.
