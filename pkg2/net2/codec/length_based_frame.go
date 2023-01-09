package codec

import (
	"bufio"
	"io"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net2"
)

func init() {
	net2.RegisterProtocol("length_based_frame", NewLengthBasedFrame(kDefaultReadBufferSize))
}

const (
	kDefaultReadBufferSize = 1024
)

func NewLengthBasedFrame(readBuf int) net2.Protocol {
	if readBuf <= 0 {
		readBuf = kDefaultReadBufferSize
	}

	return &LengthBasedFrame{
		readBuf: readBuf,
	}
}

type LengthBasedFrame struct {
	readBuf int
}

func (b *LengthBasedFrame) NewCodec(rw io.ReadWriter) (cc net2.Codec, err error) {
	codec := new(LengthBasedFrameCodec)

	codec.stream.w = rw.(io.Writer)
	codec.stream.r = bufio.NewReaderSize(rw, b.readBuf)
	codec.stream.c = rw.(io.Closer)

	return codec, nil
}

type LengthBasedFrameStream struct {
	w io.Writer
	r *bufio.Reader
	c io.Closer
}

func (s *LengthBasedFrameStream) close() error {
	if s.c != nil {
		return s.c.Close()
	}
	return nil
}

type LengthBasedFrameCodec struct {
	stream LengthBasedFrameStream
}

func (c *LengthBasedFrameCodec) Send(msg interface{}) error {
	buf := []byte(msg.(string))

	if _, err := c.stream.w.Write(buf); err != nil {
		return err
	}

	return nil
}

func (c *LengthBasedFrameCodec) Receive() (interface{}, error) {
	line, err := c.stream.r.ReadString('\n')
	return line, err
}

func (c *LengthBasedFrameCodec) Close() error {
	return c.stream.close()
}

func (c *LengthBasedFrameCodec) Context() interface{} {
	return nil
}
