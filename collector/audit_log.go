package collector

import (
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apiserver/pkg/apis/audit"
)

const Namespace = "audit_log"

type AuditLogCollector struct {
	logger micrologger.Logger

	requests *prometheus.HistogramVec
}

func New(logger micrologger.Logger) *AuditLogCollector {
	requestsOpts := prometheus.HistogramOpts{
		Namespace: Namespace,
		Name:      "requests",
		Help:      "apiserver requests processed from audit log",
	}

	labelNames := []string{"user", "user_agent"}

	return &AuditLogCollector{
		logger: logger,

		requests: prometheus.NewHistogramVec(requestsOpts, labelNames),
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

	c.requests.With(labels).Observe(event.StageTimestamp.Time.Sub(event.RequestReceivedTimestamp.Time).Seconds())
}
