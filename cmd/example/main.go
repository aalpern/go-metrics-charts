package main

import (
	_ "expvar"
	"fmt"
	"net/http"
	"time"

	"github.com/aalpern/go-metrics-charts"
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"
)

const addr = ":8080"

func main() {
	r := metrics.NewRegistry()

	metrics.RegisterDebugGCStats(r)
	metrics.RegisterRuntimeMemStats(r)
	go metrics.CaptureDebugGCStats(r, time.Second*5)
	go metrics.CaptureRuntimeMemStats(r, time.Second*5)

	exp.Exp(r)
	metricscharts.Register()

	fmt.Printf("Listening at http://localhost%s\n", addr)
	fmt.Printf("     http://localhost%s/debug/metrics\n", addr)
	fmt.Printf("     http://localhost%s/debug/metrics/charts\n", addr)

	http.ListenAndServe(":8080", nil)
}
