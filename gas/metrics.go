package gas

import (
	"github.com/go-sensors/core/gas"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	gas_concentration_count = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gas_concentration_count",
			Help: "Count of measurements of gas concentration of a given gas or group of gasses",
		},
		[]string{
			"gas",
			"source",
		},
	)
	gas_concentration_gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gas_concentration",
			Help: "Gas concentration of a given gas or group of gasses in parts per billion (ppb)",
		},
		[]string{
			"gas",
			"source",
		},
	)
)

// Set gas concentration metrics with additional labels
func SetWithLabels(concentration *gas.Concentration, labels prometheus.Labels) {
	l := map[string]string{
		"gas": concentration.Gas,
	}
	for k, v := range labels {
		l[k] = v
	}
	gas_concentration_count.With(l).Inc()
	gas_concentration_gauge.With(l).Set(concentration.Amount.PartsPerBillion())
}
