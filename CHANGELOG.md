# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.2] - 2020-12-04

### Changed

- Convert `Deployment` to `DaemonSet` and schedule POD to all masters.
- Mount `/var/log` to POD and tail `/var/log/apiserver/audit.log` for event processing.

## [0.1.0] - 2020-12-04

- First release.


[Unreleased]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.2...HEAD
[0.0.2]: https://github.com/giantswarm/k8s-audit-metrics/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.0.1
