module github.com/giantswarm/k8s-audit-metrics

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/giantswarm/exporterkit v0.2.0
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.4.0
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/prometheus/client_golang v1.8.0
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.60.1 // indirect
	k8s.io/apiserver v0.18.9
)

// We don't use etcd nor websocket but those are indirect dependencies that has
// a CVEs so put the restrictions here in order to silence nancy.
replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
)
