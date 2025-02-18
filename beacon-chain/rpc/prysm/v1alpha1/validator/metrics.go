package validator

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	errMethodNotFoundCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "execution_method_not_found_count",
		Help: "The number of errors that occurred due to method not found",
	})
)