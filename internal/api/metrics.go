package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	incSuccessCreateRecipeCounter()
	incSuccessRemoveRecipeCounter()
	incSuccessUpdateRecipeCounter()
}

type metrics struct {
	successCreateRecipeCounter prometheus.Counter
	successRemoveRecipeCounter prometheus.Counter
	successUpdateRecipeCounter prometheus.Counter
}

func newApiMetrics() Metrics {
	return &metrics{
		successCreateRecipeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_create_recipe_request_count",
			Help: "The total number of success created recipes",
		}),
		successRemoveRecipeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_remove_recipe_request_count",
			Help: "The total number of success removed recipes",
		}),
		successUpdateRecipeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_update_recipe_request_count",
			Help: "The total number of success updated recipes",
		}),
	}
}

func (m *metrics) incSuccessCreateRecipeCounter() {
	m.successCreateRecipeCounter.Inc()
}
func (m *metrics) incSuccessRemoveRecipeCounter() {
	m.successRemoveRecipeCounter.Inc()
}

func (m *metrics) incSuccessUpdateRecipeCounter() {
	m.successUpdateRecipeCounter.Inc()
}
