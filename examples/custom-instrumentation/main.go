// Package main demonstrates how to instrument httpbin with custom metrics.
package main

import (
	"net/http"

	"github.com/DataDog/datadog-go/statsd"

	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

func main() {
	statsdClient, _ := statsd.New("")

	app := httpbin.New(
		httpbin.WithObserver(datadogObserver(statsdClient)),
	)

	listenAddr := "0.0.0.0:8080"
	http.ListenAndServe(listenAddr, app)
}

func datadogObserver(client statsd.ClientInterface) httpbin.Observer {
	_ = "STUB: not implemented"
	return *new(httpbin.Observer)
}

// Log the request

// Submit a new distribution metric to datadog with tags that allow
// graphing request rate, timing, errors broken down by
// method/status/path.
