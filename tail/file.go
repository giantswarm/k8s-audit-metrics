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

	go func() {
		defer close(output)
		for {
			if err := readFile(ctx, name, output, logger); err != nil {
				logger.Errorf(ctx, err, "read")
				break
			}
		}
	}()

	return output, nil
}

func readFile(ctx context.Context, name string, output chan string, logger micrologger.Logger) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	// Seek to the end of the file to not repeat earlier lines.
	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}

	// Capture file inode number to detect log rotation.
	inodeBeginning, err := getInode(name)
	if err != nil {
		return err
	}

	reader := newLineReader(f)

	err = reader.readLines(output)
	if err != nil {
		return err
	}

	for {
		inode, err := getInode(name)
		if err != nil {
			return err
		}

		if inodeBeginning != inode {
			// File has been rotated. We need to restart.
			logger.Debugf(ctx, "inode has changed (was %q, now %q)", inodeBeginning, inode)
			break
		}

		err = reader.readLines(output)
		if err != nil {
			return err
		}

		time.Sleep(10 * time.Millisecond)
	}

	return nil
}

func getInode(fileName string) (uint64, error) {
	var s syscall.Stat_t
	err := syscall.Stat(fileName, &s)
	if err != nil {
		return 0, microerror.Mask(err)
	}

	return s.Ino, nil
}
