package server

import (
	"bufio"
	"net"
	"sync"
	"time"

	log "github.com/zeromicro/go-zero/core/logx"
)

type TcpConnWrapper struct {
	base      net.Conn
	r         *bufio.Reader
	closed    bool
	closeChan chan struct{}
	closeOnce sync.Once
	RecvChan  chan interface{}
	SendChan  chan interface{}
}

func NewTcpConnWrapper(base net.Conn) (conn *TcpConnWrapper) {
	return &TcpConnWrapper{
		base:      base,
		r:         bufio.NewReaderSize(base, 1024),
		closeChan: make(chan struct{}),
		RecvChan:  make(chan interface{}),
		SendChan:  make(chan interface{}),
	}
}

func (c *TcpConnWrapper) WrapBaseForTest(wrap func(net.Conn) net.Conn) {
	c.base = wrap(c.base)
}

func (c *TcpConnWrapper) RemoteAddr() net.Addr {
	return c.base.RemoteAddr()
}

func (c *TcpConnWrapper) LocalAddr() net.Addr {
	return c.base.LocalAddr()
}

func (c *TcpConnWrapper) SetDeadline(t time.Time) error {
	return c.base.SetDeadline(t)
}

func (c *TcpConnWrapper) SetReadDeadline(t time.Time) error {
	return c.base.SetReadDeadline(t)
}

func (c *TcpConnWrapper) SetWriteDeadline(t time.Time) error {
	return c.base.SetWriteDeadline(t)
}

func (c *TcpConnWrapper) Close() error {
	log.Info("TcpConnWrapper.Close()")
	c.closeOnce.Do(func() {
		c.closed = true
		close(c.closeChan)
		close(c.RecvChan)
		close(c.SendChan)
	})
	return c.base.Close()
}

func (c *TcpConnWrapper) Peek(n int) ([]byte, error) {
	return c.r.Peek(n)
}

func (c *TcpConnWrapper) Read(b []byte) (n int, err error) {
	n, err = c.r.Read(b)
	if err == nil {
		log.Info(string(b[:n]))
		return
	}
	c.base.Close()
	return
}

func (c *TcpConnWrapper) Write(b []byte) (n int, err error) {
	return c.base.Write(b)
}
