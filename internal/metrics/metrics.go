package metrics

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Finisher func(responseCode string)

var (
	queriesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "service_catalog",
		Name:      "query_total",
		Help:      "Count of the number of queries against the service catalog",
	}, []string{"resolver"})

	queryTimeSum = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "service_catalog",
		Name:      "query_duration_seconds_total",
		Help:      "Sum of query time broken down by resolver and status",
	}, []string{"resolver", "response_code"})
)

func NewMetricsHandler() http.Handler {
	return promhttp.Handler()
}

func RecordQuery(resolver string) Finisher {
	queriesTotal.WithLabelValues(resolver).Inc()
	startTime := time.Now()
	o := sync.Once{}
	fin := func(responseCode string) {
		endTime := time.Now()

		o.Do(func() {
			queryTimeSum.With(prometheus.Labels{"response_code": responseCode, "resolver": resolver}).Observe(endTime.Sub(startTime).Seconds())
		})
	}
	return fin
}
