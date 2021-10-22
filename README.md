[![CircleCI](https://circleci.com/gh/giantswarm/k8s-audit-metrics.svg?style=shield)](https://circleci.com/gh/giantswarm/k8s-audit-metrics) [![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/k8s-audit-metrics)](https://goreportcard.com/report/github.com/giantswarm/k8s-audit-metrics) [![](https://godoc.org/github.com/giantswarm/k8s-audit-metrics?status.svg)](http://godoc.org/github.com/giantswarm/k8s-audit-metrics) [![](https://img.shields.io/docker/pulls/giantswarm/k8s-audit-metrics.svg)](http://hub.docker.com/giantswarm/k8s-audit-metrics)

# k8s-audit-metrics

K8s-audit-metrics is a service that processes Kubernetes apiserver's audit logs and exposes metrics from it.

### Tips & tricks

#### Kubernetes client user-agent

In order to have nicer labels and easier way to distinct different clients, it's good to configure appropriate user-agent header to your k8s client.

[Client-go rest.Config](https://pkg.go.dev/k8s.io/client-go/rest#Config) has a field `UserAgent` that is useful to set to `<component>/<version>`.

Example (from [azure-operator](https://github.com/giantswarm/azure-operator/pull/1221/files)):
```
restConfig.UserAgent = fmt.Sprintf("%s/%s", project.Name(), project.Version())
```


#### Prometheus queries

##### authorization failures

Each entry in the audit log has information about authorization status and we expose that information in the metrics - `authorization_decision` tells you whether or not a request was authorized and `authorization_decision_reason` tells you why. The following query gives you the count of all requests that got forbidden:
```
count({authorization_decision="forbid"})
```

##### request duration

`k8s_api_audit_request_duration_nanoseconds` gives you information about request duration and potential latencies.


##### req/min per component

Grouping metrics by user-agent and computing rate of requests gives a metric for req/min e.g. as follows:
```
sum by (user_agent) (rate(k8s_api_audit_requests_total[5m])*60)
```

## Prerequisites

## Getting Project

Download the latest release: https://github.com/giantswarm/k8s-audit-metrics/releases/latest

Clone the git repository: https://github.com/giantswarm/k8s-audit-metrics.git

Download the latest docker image from here: https://hub.docker.com/r/giantswarm/k8s-audit-metrics/

### How to build

#### Building the standard way

```nohighlight
go build
```

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- Bugs: [issues](https://github.com/giantswarm/k8s-audit-metrics/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING.md](/giantswarm/k8s-audit-metrics/blob/master/CONTRIBUTING.md) for details on submitting patches, the contribution workflow as well as reporting bugs.

For security issues, please see [the security policy](SECURITY.md).

## License

PROJECT is under the Apache 2.0 license. See the [LICENSE](/giantswarm/k8s-audit-metrics/blob/master/LICENSE) file for details.
