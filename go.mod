module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/giantswarm/exporterkit v0.2.1
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	gopkg.in/ini.v1 v1.60.1 // indirect
	k8s.io/apiserver v0.21.1
)

// We don't use etcd nor websocket but those are indirect dependencies that has
// a CVEs so put the restrictions here in order to silence nancy.
replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
)
