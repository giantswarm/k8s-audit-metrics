# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

## [0.1.0] - 2020-12-04

- First release.


[Unreleased]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.6...HEAD
[0.0.6]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.5...v0.0.6
[0.0.5]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.0.1
