package prometheus

//A histogram for request duration, exported via a Prometheus summary with dynamically-computed quantiles.

import (
	"time"

	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	var dur metrics.Histogram = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "myservice",
		Subsystem: "api",
		Name:     "request_duration_seconds",
		Help:     "Total time spent serving requests.",
	}, []string{})
	// ...
}

func handleRequest(dur metrics.Histogram) {
	defer func(begin time.Time) { dur.Observe(time.Since(begin).Seconds()) }(time.Now())
	// handle request
}
