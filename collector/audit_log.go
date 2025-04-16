package collector

import (
	"regexp"

	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apiserver/pkg/apis/audit"
)

const (
	Namespace = "k8s_api_audit"

	// https://github.com/kubernetes/apiserver/blob/706a6d89cf35950281e095bb1eeed5e3211d6272/pkg/endpoints/filters/authorization.go#L35-L36
	decisionAnnotationKey = "authorization.k8s.io/decision"
	reasonAnnotationKey   = "authorization.k8s.io/reason"
	// https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1205-bound-service-account-tokens/README.md
	staleTokenAnnotationKey = "authentication.k8s.io/stale-token"
)

var metricLabels = []string{
	"authorization_decision",
	"authorization_decision_reason",
	"authentication_stale_token",
	"user",
	"user_agent",
}

type AuditLogCollector struct {
	logger micrologger.Logger

	requestsCount    *prometheus.CounterVec
	requestsDuration *prometheus.GaugeVec

	staleTokenRegex *regexp.Regexp
}

func New(logger micrologger.Logger) *AuditLogCollector {
	return &AuditLogCollector{
		logger: logger,

		staleTokenRegex: regexp.MustCompile(`^subject: (?P<serviceAccount>.+), seconds after warning threshold: (?P<threshold>.+)$`),

		requestsCount:    newRequestCountMetric(),
		requestsDuration: newRequestDurationMetric(),
	}
}

func (c *AuditLogCollector) Describe(ch chan<- *prometheus.Desc) {
	c.requestsCount.Describe(ch)
	c.requestsDuration.Describe(ch)
}

func (c *AuditLogCollector) Collect(ch chan<- prometheus.Metric) {
	c.requestsCount.Collect(ch)
	c.requestsDuration.Collect(ch)
}

func (c *AuditLogCollector) Process(event audit.Event) {
	prometheusLabels := c.buildMetricLabels(event)
	c.requestsCount.With(prometheusLabels).Inc()

	if event.Stage == audit.StageResponseComplete {
		duration := event.StageTimestamp.Sub(event.RequestReceivedTimestamp.Time)
		c.requestsDuration.With(prometheusLabels).Set(float64(duration))
	}
}

func (c *AuditLogCollector) buildMetricLabels(event audit.Event) prometheus.Labels {
	staleServiceAccountToken := ""
	if event.Annotations[staleTokenAnnotationKey] != "" {
		res := c.staleTokenRegex.FindStringSubmatch(event.Annotations[staleTokenAnnotationKey])
		staleServiceAccountToken = res[1]
	}

	return prometheus.Labels{
		"authorization_decision":        event.Annotations[decisionAnnotationKey],
		"authorization_decision_reason": event.Annotations[reasonAnnotationKey],
		"authentication_stale_token":    staleServiceAccountToken,
		"user":                          event.User.Username,
		"user_agent":                    event.UserAgent,
	}
}

func newRequestCountMetric() *prometheus.CounterVec {
	opts := prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "requests_total",
		Help:      "apiserver requests count processed from audit log",
	}

	return prometheus.NewCounterVec(opts, metricLabels)
}

func newRequestDurationMetric() *prometheus.GaugeVec {
	opts := prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "request_duration_nanoseconds",
		Help:      "apiserver requests duration processed from audit log",
	}

	return prometheus.NewGaugeVec(opts, metricLabels)
}
