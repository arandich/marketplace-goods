package metrics

import (
	"github.com/arandich/marketplace-goods/internal/config"
	sdkPrometheus "github.com/arandich/marketplace-sdk/prometheus"
)

type Metrics struct {
	// Base metrics.
	BaseMetrics sdkPrometheus.Metrics
}

func New(baseMetrics sdkPrometheus.Metrics, cfg config.PrometheusConfig) Metrics {
	return Metrics{BaseMetrics: baseMetrics}
}
