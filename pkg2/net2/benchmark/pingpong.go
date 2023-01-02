package main

import (
	"io"
	"net"

	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net2"
)

func init() {
	net2.RegisterProtocol("pingpong", NewPingpong(4096))
}

type Pingpong struct {
	readBuf int
}

func NewPingpong(readBuf int) net2.Protocol {
	if readBuf <= 0 {
		readBuf = 4096
	}

	return &Pingpong{
		readBuf: readBuf,
	}
}

func (b *Pingpong) NewCodec(rw io.ReadWriter) (cc net2.Codec, err error) {
	codec := new(PingpongCodec)

	codec.readBuf = b.readBuf
	codec.rw = rw
	codec.c = rw.(io.Closer)

	codec.buf = make([]byte, b.readBuf)
	return codec, nil
}

type PingpongCodec struct {
	readBuf int
	rw      io.ReadWriter
	c       io.Closer
	buf     []byte
}

func (c *PingpongCodec) Send(msg interface{}) error {
	if _, err := c.rw.Write(msg.([]byte)); err != nil {
		return err
	}
	return nil
}

func (c *PingpongCodec) Receive() (interface{}, error) {
	buf := make([]byte, c.readBuf)
	n, err := c.rw.Read(buf)
	if err == nil {
		return buf[:n], nil
	}
	return nil, err
}

func (c *PingpongCodec) Close() error {
	return c.c.Close()
}

func (c *PingpongCodec) Context() interface{} {
	return nil
}

// PingpongServer
// ///////////////////////////////////////////////////////////////////////////////////////
type PingpongServer struct {
	server *net2.TcpServer
}

func NewPingpongServer(listener net.Listener, protoName string) *PingpongServer {
	s := &PingpongServer{}
	s.server = net2.NewTcpServer(net2.TcpServerArgs{Listener: listener, ServerName: "pingpong", ProtoName: protoName, SendChanSize: 0, ConnectionCallback: s, MaxConcurrentConnection: 2})
	return s
}

func (s *PingpongServer) Serve() {
	s.server.Serve()
}

func (s *PingpongServer) OnNewConnection(conn *net2.TcpConnection) {
	logx.Info("OnNewConnection %v", conn.RemoteAddr())
}

func (s *PingpongServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	// log.Infof("echo_server recv peer(%v) data: %v", conn.RemoteAddr(), msg)
	conn.Send(msg)
	return nil
}

func (s *PingpongServer) OnConnectionClosed(conn *net2.TcpConnection) {
	logx.Info("OnConnectionClosed - %v", conn.RemoteAddr())
}

// ///////////////////////////////////////////////////////////////////////////////////////
type PingpongInsance struct {
	server *PingpongServer
}

func (this *PingpongInsance) Initialize() error {
	listener, err := net.Listen("tcp", "0.0.0.0:33333")
	if err != nil {
		logx.Error("listen error: %v", err)
		return err
	}

	this.server = NewPingpongServer(listener, "pingpong")
	return nil
}

func (this *PingpongInsance) RunLoop() {
	this.server.Serve()
}

func (this *PingpongInsance) Destroy() {
	this.server.server.Stop()
}

func main() {
	// instance := &PingpongInsance{}
	// util.DoMainAppInstance(instance)
}
