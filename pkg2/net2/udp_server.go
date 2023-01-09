package net2

import (
	"net"
	"sync"

	log "github.com/zeromicro/go-zero/core/logx"
)

//var ConnectionClosedError = errors.New("Connection Closed")
//var ConnectionBlockedError = errors.New("Connection Blocked")

type UdpConnectionCallback interface {
	OnUdpNew(conn *UdpConnection)
	OnUdpDataArrived(c *UdpConnection, msg interface{}) error
	OnUdpClosed(c *UdpConnection)
}

type UdpServer struct {
	conn         *net.UDPConn
	serverName   string
	protoName    string
	sendChanSize int
	running      bool
	sem          chan struct{}
	releaseOnce  sync.Once
	conns        map[string]*UdpConnection
}

type UdpServerArgs struct {
	PacketConn   *net.UDPConn
	ServerName   string
	ProtoName    string
	SendChanSize int
}

func NewUdpServer(args UdpServerArgs) *UdpServer {
	return &UdpServer{
		conn:         args.PacketConn,
		serverName:   args.ServerName,
		protoName:    args.ProtoName,
		sendChanSize: args.SendChanSize,
		running:      false,
	}
}

func ListenUdp(addr2 string) (net.PacketConn, error) {
	//start tcp server
	addr, err := net.ResolveUDPAddr("udp", addr2)
	if err != nil {
		log.Error("net.ResolveUDPAddr fail - %v", err)
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Error("net.ListenUDP fail - %v", err)
		//os.Exit(1)
		return nil, err
	}
	return conn, nil
}

func (s *UdpServer) getOrCreateConn(addr *net.UDPAddr) *UdpConnection {
	a := addr.String()
	if c, ok := s.conns[a]; ok {
		return c
	} else {
		c = NewUdpConnection()
		s.conns[a] = c
		return c
	}
}

func (s *UdpServer) Serve() {
	if s.running {
		return
	}
	s.running = true
	s.acquire()

	buf := make([]byte, 0, 4096)
	for {
		bufLen, remote, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Error("readFromUDP error - %v", err)
			break
		}
		_ = bufLen
		_ = remote
	}

	s.running = false
}

func (s *UdpServer) Stop() {
	if s.running {
		s.conn.Close()
		s.releaseOnce.Do(s.release)
	}
}

func (s *UdpServer) Pause() {
}

func (s *UdpServer) acquire() {
	s.sem <- struct{}{}
}

func (s *UdpServer) release() {
	<-s.sem
}
