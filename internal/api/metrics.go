package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	incSuccessCreateRecipeCounter()
	incFailCreateRecipeCounter()
	incSuccessRemoveRecipeCounter()
	incFailRemoveRecipeCounter()
	incSuccessUpdateRecipeCounter()
	incFailUpdateRecipeCounter()
}

const (
	successResultLabel = "success"
	failResultLabel = "fail"
)
var labels = []string{"result"}

type metrics struct {
	createRecipeCounter *prometheus.CounterVec
	removeRecipeCounter *prometheus.CounterVec
	updateRecipeCounter *prometheus.CounterVec
}

func newApiMetrics() Metrics {
	return &metrics{
		createRecipeCounter: promauto.NewCounterVec(prometheus.CounterOpts{
				Name: "create_recipe_request_count",
				Help: "number of created recipes",
			},
			labels),

		removeRecipeCounter: promauto.NewCounterVec(prometheus.CounterOpts{
				Name: "remove_recipe_request_count",
				Help: "number of removed recipes",
			},
			labels),
		updateRecipeCounter: promauto.NewCounterVec(prometheus.CounterOpts{
				Name: "update_recipe_request_count",
				Help: "number of removed recipes",
			},
			labels),
	}
}

func (m *metrics) incSuccessCreateRecipeCounter() {
	m.createRecipeCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailCreateRecipeCounter() {
	m.createRecipeCounter.WithLabelValues(failResultLabel).Inc()
}

func (m *metrics) incSuccessRemoveRecipeCounter() {
	m.removeRecipeCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailRemoveRecipeCounter() {
	m.removeRecipeCounter.WithLabelValues(failResultLabel).Inc()
}

func (m *metrics) incSuccessUpdateRecipeCounter() {
	m.updateRecipeCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailUpdateRecipeCounter() {
	m.updateRecipeCounter.WithLabelValues(failResultLabel).Inc()
}