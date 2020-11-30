module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/micrologger v0.3.4
	github.com/google/go-cmp v0.5.3 // indirect
)

replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
