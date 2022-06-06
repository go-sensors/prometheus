package temperature

import (
	"github.com/go-sensors/core/units"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	temperature_count = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "temperature_count",
			Help: "Count of measurements of temperature",
		},
		[]string{
			"source",
		},
	)
	temperature_gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature",
			Help: "Temperature in degrees Celsius (°C)",
		},
		[]string{
			"source",
		},
	)
)

// Set temperature metrics with additional labels
func SetWithLabels(temperature *units.Temperature, labels prometheus.Labels) {
	temperature_count.With(labels).Inc()
	temperature_gauge.With(labels).Set(temperature.DegreesCelsius())
}
