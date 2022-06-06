package pm

import (
	"github.com/go-sensors/core/pm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	pm_concentration_count = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pm_concentration_count",
			Help: "Count of measurements of particulate matter concentration of a given upper bound size",
		},
		[]string{
			"upper_bound_size",
			"source",
		},
	)
	pm_concentration_gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pm_concentration",
			Help: "Particulate matter concentration of a given upper bound size in micrograms per cubic meter (µg/㎥)",
		},
		[]string{
			"upper_bound_size",
			"source",
		},
	)
)

// Set particulate matter concentration metrics with additional labels
func SetWithLabels(concentration *pm.Concentration, labels prometheus.Labels) {
	l := map[string]string{
		"upper_bound_size": concentration.UpperBoundSize.String(),
	}
	for k, v := range labels {
		l[k] = v
	}
	pm_concentration_count.With(l).Inc()
	pm_concentration_gauge.With(l).Set(concentration.Amount.MicrogramsPerCubicMeter())
}
