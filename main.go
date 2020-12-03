package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/giantswarm/exporterkit"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apiserver/pkg/apis/audit"

	"github.com/giantswarm/k8s-audit-metrics/collector"
	"github.com/giantswarm/k8s-audit-metrics/tail"
)

func main() {
	err := mainError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", microerror.JSON(err)))
	}
}

func mainError() error {
	logger, err := micrologger.New(micrologger.Config{})
	if err != nil {
		return microerror.Mask(err)
	}

	var lines <-chan string
	if len(os.Args) > 1 {
		lines, err = tail.File(os.Args[1], logger)
		if err != nil {
			return microerror.Mask(err)
		}
	} else {
		lines, err = tail.Stdin(logger)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	auditCollector := collector.New(logger)

	err = startExporterServer(logger, auditCollector)
	if err != nil {
		return microerror.Mask(err)
	}

	err = run(logger, lines, auditCollector)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func run(log micrologger.Logger, lines <-chan string, metrics *collector.AuditLogCollector) error {
	for line := range lines {
		var event audit.Event
		err := json.Unmarshal([]byte(line), &event)
		if err != nil {
			log.Errorf(context.Background(), err, "json.Unmarshal()")
			continue
		}

		metrics.Process(event)
	}

	return nil
}

func startExporterServer(logger micrologger.Logger, collectors ...prometheus.Collector) error {
	c := exporterkit.Config{
		Collectors: collectors,
		Logger:     logger,
	}

	exporter, err := exporterkit.New(c)
	if err != nil {
		return microerror.Mask(err)
	}

	go exporter.Run()

	return nil
}
