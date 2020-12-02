package tail

import (
	"bufio"
	"io"

	"github.com/giantswarm/microerror"
)

type lineReader struct {
	line   []byte
	reader *bufio.Reader
}

func newLineReader(r io.Reader) *lineReader {
	return &lineReader{
		reader: bufio.NewReader(r),
	}
}

func (r *lineReader) readLines(output chan<- string) error {
	var err error
	var bs []byte

	for {
		bs, err = r.reader.ReadBytes('\n')
		r.line = append(r.line, bs...)

		if len(r.line) > 0 && hasEOL(r.line) {
			output <- string(trimEOL(r.line))
			r.line = r.line[:0]
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func hasEOL(bs []byte) bool {
	n := len(bs)

	if bs[n-1] == '\n' || bs[n-1] == '\r' {
		return true
	}

	return false
}

func trimEOL(bs []byte) []byte {
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == '\n' || bs[i] == '\r' {
			bs = bs[:i]
		}
	}

	return bs
}
