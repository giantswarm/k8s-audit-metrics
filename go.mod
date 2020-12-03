module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/exporterkit v0.2.0
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.4.0
	github.com/prometheus/client_golang v1.7.1
	k8s.io/apiserver v0.18.9
)

// We don't use websocket but it's indirect dependency that fixes a CVE so put
// the definition here in order to silence nancy.
replace github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
