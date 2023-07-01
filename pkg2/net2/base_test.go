package net2

/*
import (
	"bufio"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestCodec struct {
	*bufio.Reader
	io.Writer
	io.Closer
	mt string
}

func (c *TestCodec) Send(msg interface{}) error {
	buf := []byte(msg.(string))
	if _, err := c.Writer.Write(buf); err != nil {
		return err
	}

	return nil
}

func (c *TestCodec) Receive() (interface{}, error) {
	line, err := c.Reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return line, err
}

func (c *TestCodec) Close() error {
	return c.Closer.Close()
}

func (c *TestCodec) ClearSendChan(ic <-chan interface{}) {
	logx.Info("TestCodec ClearSendChan, %v", ic)
}

// ////////////////////////////////////////////////////////////////////////////////////////
type TestProto struct {
}

func (b *TestProto) NewCodec(rw io.ReadWriter) (cc Codec, err error) {
	c := &TestCodec{
		Reader: bufio.NewReader(rw),
		Writer: rw.(io.Writer),
		Closer: rw.(io.Closer),
	}
	return c, nil
}
*/
