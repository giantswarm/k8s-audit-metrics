# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Update Kyverno `PolicyExceptions` to `v2beta1`.

## [0.9.0] - 2024-04-01

### Added

- Add team label in resources.
- Use ServiceMonitor for monitoring.

### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.

## [0.8.0] - 2023-10-18

### Changed

- Replace condition for PSP CR installation.

## [0.7.1] - 2023-09-19

### Changed

- Removed `/metrics` checks in cilium network policy

## [0.7.0] - 2023-09-13

### Changed

- Switched to `kube-system` namespace by default
- Added Cilium Network Policy to scrape `/metrics` on port `8000`

## [0.6.1] - 2023-08-03

### Changed

- Push to default app catalog.

## [0.6.0] - 2023-07-13

### Fixed

- Added required values for pss policies.
- Added pss exceptions.

## [0.5.3] - 2023-03-22

### Changed

- Reduce cardinality for the `authentication_stale_token` label.

## [0.5.2] - 2022-11-02

### Changed

- Added option to set pod's `priorityClassName`.

## [0.5.1] - 2022-10-25

### Changed

- Push to Aliyun.

## [0.5.0] - 2022-09-14

### Changed

- Change helm chart's value path for the image registry.
- Push app to Azure app collection.

## [0.4.2] - 2022-09-12

### Fixed

- Fix PSP.

## [0.4.1] - 2022-09-12

### Changed

- Run as root to avoid permission errors.

## [0.4.0] - 2022-09-12

### Changed

- Push app to AWS app collection.

## [0.3.0] - 2022-03-21

### Added

- Add VerticalPodAutoscaler CR.

## [0.2.0] - 2021-12-14

### Added
- Adds new authentication.k8s.io/stale-token audit annotation, introduced by Bound Service Account tokens enhancement, as metric label.

## [0.1.0] - 2020-12-04

- First release.

## [0.0.7] - 2020-12-07

### Changed

- Quote monitoring port value in service annotation.

## [0.0.6] - 2020-12-07

### Added

- Added missing annotations to get service picked up from tenant cluster by prometheus.

## [0.0.5] - 2020-12-07

### Changed

- Push app to CP app catalog as well.

## [0.0.4] - 2020-12-04

- Change `Service` type to default.
- Label `Service` with `giantswarm.io/monitoring` to enable scraping.

## [0.0.3] - 2020-12-04

### Changed

- Change request metric from latency histogram to simple counter.

## [0.0.2] - 2020-12-04

### Changed

- Convert `Deployment` to `DaemonSet` and schedule POD to all masters.
- Mount `/var/log` to POD and tail `/var/log/apiserver/audit.log` for event processing.

## [0.0.2] - 2020-12-04

- First release.


[Unreleased]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.9.0...HEAD
[0.9.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.6.1...v0.7.0
[0.6.1]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.6.0...v0.6.0
[0.6.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.6.0...v0.6.0
[0.6.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.6.0...v0.6.0
[0.6.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.5.3...v0.6.0
[0.5.3]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.5.2...v0.5.3
[0.5.2]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.5.1...v0.5.2
[0.5.1]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.4.2...v0.5.0
[0.4.2]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.4.1...v0.4.2
[0.4.1]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.7...v0.1.0
[0.0.7]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.6...v0.0.7
[0.0.6]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.5...v0.0.6
[0.0.5]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.0.1
