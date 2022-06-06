package humidity

import (
	"github.com/go-sensors/core/units"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	relative_humidity_count = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "relative_humidity_count",
			Help: "Count of measurements of relative humidity",
		},
		[]string{
			"source",
		},
	)
	relative_humidity_gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "relative_humidity",
			Help: "Relative humidity (typical range 0.0 .. 1.0)",
		},
		[]string{
			"source",
		},
	)
	absolute_humidity_gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "absolute_humidity",
			Help: "Absolute humidity in micrograms per cubic meter (µg/㎥)",
		},
		[]string{
			"source",
		},
	)
)

// Set relative humidity metrics with additional labels
func SetWithLabels(relativeHumidity *units.RelativeHumidity, labels prometheus.Labels) {
	relative_humidity_count.With(labels).Inc()
	relative_humidity_gauge.With(labels).Set(relativeHumidity.Percentage)
	absolute_humidity_gauge.With(labels).Set(relativeHumidity.AbsoluteHumidity().MicrogramsPerCubicMeter())
}
