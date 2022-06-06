package stat

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Counter single value with counter
type Counter struct {
	name  string
	t     time.Time
	gauge prometheus.Gauge
}

func newCounter(name string) *Counter {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
	})
	prometheus.MustRegister(gauge)
	return &Counter{
		name:  name,
		t:     time.Now(),
		gauge: gauge,
	}
}

// Set set counter value
func (ct *Counter) Set(v float64) {
	ct.gauge.Set(v)
}

// Inc increase counter value
func (ct *Counter) Inc() {
	ct.gauge.Inc()
}

// Dec decrease counter value
func (ct *Counter) Dec() {
	ct.gauge.Dec()
}
