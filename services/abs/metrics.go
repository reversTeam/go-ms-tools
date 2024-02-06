package abs

import (
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promauto"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	ServiceName string
	// requests *prometheus.GaugeVec
}

// requests: promauto.NewGaugeVec(
// prometheus.GaugeOpts{
// 	Name: "fizzbuzz_request_sec",
// 	Help: "Number of requests",
// }, []string{"code", "method", "path"}),

func NewMetrics(serviceName string) *Metrics {
	return &Metrics{
		ServiceName: serviceName,
	}
}

func (o *Metrics) GetServiceName() string {
	return o.ServiceName
}
