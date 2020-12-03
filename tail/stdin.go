package tail

import (
	"context"
	"os"

	"github.com/giantswarm/micrologger"
)

func Stdin(logger micrologger.Logger) (<-chan string, error) {
	output := make(chan string)

	go func() {
		reader := newLineReader(os.Stdin)
		err := reader.readLines(output)
		if err != nil {
			logger.Errorf(context.Background(), err, "read failed")
		}

		close(output)
	}()

	return output, nil
}
