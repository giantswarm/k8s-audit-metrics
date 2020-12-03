package tail

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

func File(name string, logger micrologger.Logger) (<-chan string, error) {
	ctx := context.Background()
	output := make(chan string, 2)

	f, err := os.Open(name)
	if err != nil {
		logger.Errorf(ctx, err, "open")
		return nil, microerror.Mask(err)
	}

	inodeBeginning, err := getInode(name)
	if err != nil {
		logger.Errorf(ctx, err, "stat")
		return nil, microerror.Mask(err)
	}

	reader := newLineReader(f)

	err = reader.readLines(output)
	if err != nil {
		logger.Errorf(ctx, err, "read")
		return nil, microerror.Mask(err)
	}

	go func() {
		for {
			inode, err := getInode(name)
			if err != nil {
				logger.Errorf(ctx, err, "stat")
				break
			}

			if inodeBeginning != inode {
				// File has been rotated. We need to restart.
				logger.Debugf(ctx, "inode has changed (was %q, now %q)", inodeBeginning, inode)
				break
			}

			err = reader.readLines(output)
			if err != nil {
				logger.Errorf(ctx, err, "read")
				break
			}

			time.Sleep(100 * time.Millisecond)
		}

		f.Close()
		close(output)
	}()

	return output, nil
}

func getInode(fileName string) (uint64, error) {
	var s syscall.Stat_t
	err := syscall.Stat(fileName, &s)
	if err != nil {
		return 0, microerror.Mask(err)
	}

	return s.Ino, nil
}
