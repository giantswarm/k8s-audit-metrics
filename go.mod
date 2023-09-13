module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/giantswarm/exporterkit v1.0.0
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v1.0.0
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/spf13/viper v1.16.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	k8s.io/apiserver v0.28.1
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
)

// We don't use etcd nor websocket but those are indirect dependencies that has
// a CVEs so put the restrictions here in order to silence nancy.
replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
)
