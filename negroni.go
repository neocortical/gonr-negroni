package gonroni

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/neocortical/gonr"
	"github.com/neocortical/nrmetrics"
	metrics "github.com/rcrowley/go-metrics"
)

func AddNegroniHttpMetrics(agent gonr.GonrAgent) negroni.HandlerFunc {
	responseTimer := metrics.NewTimer()

	nrmetrics.AddTimerMetric(agent.Plugin(), responseTimer, nrmetrics.MetricConfig{
		Name:        "HTTP/Response Time",
		Unit:        "requests",
		Duration:    time.Millisecond,
		Min:         true,
		Max:         true,
		Mean:        true,
		Percentiles: []float64{0.5, 0.75, 0.9, 0.99, 0.999},
		Rate1:       true,
		Rate5:       true,
		Rate15:      true,
	})

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		t := time.Now()

		next.ServeHTTP(rw, r)

		responseTimer.UpdateSince(t)
	}
}
