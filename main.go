package main

import (
	"atlas/config"
	"atlas/inspect"
	"atlas/metrics"
	"atlas/proxy"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config := config.NewConfig()
	prometheus.MustRegister(
		metrics.RequestCounter,
		metrics.RequestBlockedCounter,
		metrics.RequestFailedCounter,
	)

	inspect := inspect.NewInspectHTTPRequest(config.DenyIPList, config.DenyHTTPHeader, config.DenyHTTPBody)
	proxy := proxy.NewProxy(config.Backend, inspect)

	http.HandleFunc("/", proxy.Server)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting WaF atlas service..")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
