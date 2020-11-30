module github.com/giantswarm/k8s-audit-metrics

go 1.15

replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs

require (
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.4
)
