package common

import (
	"io"

	"github.com/golang/snappy"
	"github.com/lucas-clemente/quic-go"
)

// CompressedStream is a compressed version of
// quic stream which compresses data on write
// and de-compresses data on read
type CompressedStream struct {
	quic.Stream
	rd io.Reader
	wr io.Writer
}

// NewCompressedStream returns a new CompressedStream instance
// by initializing snappy compression methods for read and write methods
func NewCompressedStream(stream quic.Stream) *CompressedStream {
	return &CompressedStream{
		Stream: stream,
		rd:     snappy.NewReader(stream),
		wr:     snappy.NewWriter(stream),
	}
}

// Read implements the io.Reader interface
func (c *CompressedStream) Read(p []byte) (int, error) {
	return c.Stream.Read(p)
}

// Read implements the io.Writer interface
func (c *CompressedStream) Write(p []byte) (int, error) {
	return c.Stream.Write(p)
}
