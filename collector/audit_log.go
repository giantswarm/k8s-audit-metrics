package collector

import (
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apiserver/pkg/apis/audit"
)

const Namespace = "k8s_api_audit"

type AuditLogCollector struct {
	logger micrologger.Logger

	requests *prometheus.CounterVec
}

func New(logger micrologger.Logger) *AuditLogCollector {
	requestsOpts := prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "requests_total",
		Help:      "apiserver requests processed from audit log",
	}

	labelNames := []string{"user", "user_agent"}

	return &AuditLogCollector{
		logger: logger,

		requests: prometheus.NewCounterVec(requestsOpts, labelNames),
	}
}

func (c *AuditLogCollector) Describe(ch chan<- *prometheus.Desc) {
	c.requests.Describe(ch)
}

func (c *AuditLogCollector) Collect(ch chan<- prometheus.Metric) {
	c.requests.Collect(ch)
}

func (c *AuditLogCollector) Process(event audit.Event) {
	labels := prometheus.Labels{
		"user":       event.User.Username,
		"user_agent": event.UserAgent,
	}

	c.requests.With(labels).Inc()
}
