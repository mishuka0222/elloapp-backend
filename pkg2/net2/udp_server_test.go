package net2

import (
	"net"
)

type testUdpEchoServer struct {
	server      *UdpServer
	serverName  string
	workLoadCnt int
}

func NewUdpTestServer(listener net.Listener, serverName, protoName string, chanSize int, maxConn int) *testUdpEchoServer {
	s := &testUdpEchoServer{}
	s.server = NewUdpServer(
		UdpServerArgs{
			ServerName:   serverName,
			ProtoName:    protoName,
			SendChanSize: chanSize,
		})
	s.serverName = serverName
	s.workLoadCnt = 0
	return s
}

func (s *testUdpEchoServer) Serve() {
	s.server.Serve()
}

func (s *testUdpEchoServer) Stop() {
	s.server.Stop()
}
