module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.4.0
)

// We don't use websocket but it's indirect dependency that fixes a CVE so put
// the definition here in order to silence nancy.
replace github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2
