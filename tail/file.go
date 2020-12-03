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

	// Seek to the end of the file to not repeat earlier lines.
	_, err = f.Seek(0, 2)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	// Capture file inode number to detect log rotation.
	inodeBeginning, err := getInode(name)
	if err != nil {
		logger.Errorf(ctx, err, "stat")
		return nil, microerror.Mask(err)
	}

	reader := newLineReader(f)

	go func() {
		defer f.Close()
		defer close(output)

		err = reader.readLines(output)
		if err != nil {
			logger.Errorf(ctx, err, "read")
			return
		}

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

			time.Sleep(10 * time.Millisecond)
		}
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
