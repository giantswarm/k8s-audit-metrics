package main

import (
	"fmt"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/k8s-audit-metrics/pkg/logger"
)

func main() {
	err := mainError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", err))
	}
}

func mainError() error {
	var newLogger logger.Logger
	{
		mLogger, err := micrologger.New(micrologger.Config{})
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger = logger.Wrap(mLogger)
	}

	newLogger.Info("k8s-audit-metrics")

	return nil
}
