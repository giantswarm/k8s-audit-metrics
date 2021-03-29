FROM alpine:3.13.3

RUN apk add --no-cache ca-certificates

ADD ./k8s-audit-metrics /k8s-audit-metrics

ENTRYPOINT ["/k8s-audit-metrics"]
