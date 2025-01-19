package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "number_of_requests",
		Help: "The number of the HTTP requests",
	},
)

var RequestBlockedCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "number_of_blocked_requests",
		Help: "The number of the blocked HTTP requests",
	},
)

var RequestFailedCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "number_of_failed_requests",
		Help: "The number of the failed HTTP requests",
	},
)
