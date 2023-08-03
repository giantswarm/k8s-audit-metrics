module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/exporterkit v1.0.0
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v1.0.0
	github.com/prometheus/client_golang v1.12.1
	k8s.io/apiserver v0.20.15
)

// We don't use etcd nor websocket but those are indirect dependencies that has
// a CVEs so put the restrictions here in order to silence nancy.
replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
)
