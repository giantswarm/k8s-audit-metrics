module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.4
)

// We do not directly use the websocket package but within the dependency graph
// this package is necessary. We have to make sure it is at least at v1.4.2 due
// to some security fixes that CI would complain about otherwise.
replace github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
