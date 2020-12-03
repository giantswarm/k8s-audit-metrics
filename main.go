package main

import (
	"fmt"
	"log"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

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

	for s := range lines {
		fmt.Printf(": %q\n", s)
	}

	log.Println("<<<< EXIT >>>>")

	return nil
}
