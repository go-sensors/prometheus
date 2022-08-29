package prometheus

import (
	"context"

	coregas "github.com/go-sensors/core/gas"
	corepm "github.com/go-sensors/core/pm"
	"github.com/go-sensors/core/units"
	"github.com/go-sensors/prometheus/gas"
	"github.com/go-sensors/prometheus/humidity"
	"github.com/go-sensors/prometheus/pm"
	"github.com/go-sensors/prometheus/temperature"
)

const (
	// DefaultSource is the default value for the `source` Prometheus label
	DefaultSource string = "default"
)

var (
	defaultLabels = &Labels{
		Source: DefaultSource,
	}
)

// Labels defines additional labels that may be applied when setting metrics
type Labels struct {
	// Source is a user-defined string that identifies the source of the measurement
	Source string
}

// MetricHandler is an implementation of handlers for each supported type of sensor data
type MetricHandler struct {
	labels map[string]string
}

// NewMetricHandler returns a new MetricHandler with default labels
func NewMetricHandler() *MetricHandler {
	return NewMetricHandlerWithLabels(defaultLabels)
}

// NewMetricHandlerWithLabels returns a new MetricHandler with given labels
func NewMetricHandlerWithLabels(labels *Labels) *MetricHandler {
	return &MetricHandler{
		map[string]string{
			"source": labels.Source,
		},
	}
}

// Handles an individual concentration measurement
func (h *MetricHandler) HandleGasConcentration(_ context.Context, g *coregas.Concentration) error {
	gas.SetWithLabels(g, h.labels)
	return nil
}

// Handles an individual concentration measurement
func (h *MetricHandler) HandleRelativeHumidity(_ context.Context, rh *units.RelativeHumidity) error {
	humidity.SetWithLabels(rh, h.labels)
	return nil
}

// Handles an individual concentration measurement
func (h *MetricHandler) HandlePMConcentration(_ context.Context, c *corepm.Concentration) error {
	pm.SetWithLabels(c, h.labels)
	return nil
}

// Handles an individual temperature measurement
func (h *MetricHandler) HandleTemperature(_ context.Context, t *units.Temperature) error {
	temperature.SetWithLabels(t, h.labels)
	return nil
}
