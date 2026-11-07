FROM alpine:3.24.4

RUN apk add --no-cache ca-certificates

ARG TARGETARCH
COPY k8s-audit-metrics-linux-${TARGETARCH} /k8s-audit-metrics

ENTRYPOINT ["/k8s-audit-metrics"]
