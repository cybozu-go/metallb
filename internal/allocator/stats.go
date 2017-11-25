package allocator

import "github.com/prometheus/client_golang/prometheus"

var stats = struct {
	poolCapacity *prometheus.GaugeVec
	poolActive   *prometheus.GaugeVec
}{
	poolCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "metallb",
		Subsystem: "allocator",
		Name:      "addresses_total",
		Help:      "Number of usable IP addresses, per pool",
	}, []string{
		"pool",
	}),
	poolActive: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "metallb",
		Subsystem: "allocator",
		Name:      "addresses_in_use_total",
		Help:      "Number of IP addresses in use, per pool",
	}, []string{
		"pool",
	}),
}

func init() {
	prometheus.MustRegister(stats.poolCapacity)
	prometheus.MustRegister(stats.poolActive)
}
