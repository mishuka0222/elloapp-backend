package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"sync"

	log "github.com/zeromicro/go-zero/core/logx"
)

// /////////////////////////////////////////////////////////////////////////////////
var _ net.Listener = &HttpListener{}

type HttpListener struct {
	base       net.Listener
	acceptChan chan net.Conn
	closed     bool
	closeOnce  sync.Once
	closeChan  chan struct{}
}

func (l *HttpListener) Addr() net.Addr {
	return l.base.Addr()
}

func (l *HttpListener) Close() error {
	log.Info("HttpListener.Close()")
	l.closeOnce.Do(func() {
		l.closed = true
		close(l.closeChan)
	})
	return l.base.Close()
}

func (l *HttpListener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.acceptChan:
		return conn, nil
	case <-l.closeChan:
	}
	return nil, os.ErrInvalid
}

// /////////////////////////////////////////////////////////////////////////////////
func onMTProtoHttpApiw1(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)

	connPtr := writerToConnPtr(w)
	connMutex.Lock()
	defer connMutex.Unlock()

	// Requests can access connections by pointer from the responseWriter object
	conn, ok := conns[connPtr]
	if !ok {
		log.Info("error: no matching connection found")
		return
	} else {
		// defer conn.Close()
		conn.(*TcpConnWrapper).RecvChan <- req
		// _ = conn
		// _, _ = w.Write([]byte(req.RequestURI + "\n"))
		msgData, _ := <-conn.(*TcpConnWrapper).SendChan
		log.Infof("%v", msgData)
		w.Write([]byte(msgData.(*http.Request).RequestURI + "\n"))
		// close(conn.(*TcpConnWrapper).RecvChan)
	}
}

// Connection array indexed by connection address
var conns = make(map[uintptr]net.Conn)
var connMutex = sync.Mutex{}

// writerToConnPrt converts an http.ResponseWriter to a pointer for indexing
func writerToConnPtr(w http.ResponseWriter) uintptr {
	ptrVal := reflect.ValueOf(w)
	val := reflect.Indirect(ptrVal)

	// http.conn
	valconn := val.FieldByName("conn")
	val1 := reflect.Indirect(valconn)

	// net.TCPConn
	ptrRwc := val1.FieldByName("rwc").Elem()
	rwc := reflect.Indirect(ptrRwc)

	// net.Conn
	val1conn := rwc.FieldByName("conn")
	val2 := reflect.Indirect(val1conn)

	return val2.Addr().Pointer()
}

// connToPtr converts a net.Conn into a pointer for indexing
func connToPtr(c net.Conn) uintptr {
	ptrVal := reflect.ValueOf(c)
	return ptrVal.Pointer()
}

// ConnStateListener bound to server and maintains a list of connections by pointer
func ConnStateListener(c net.Conn, cs http.ConnState) {
	connPtr := connToPtr(c)
	connMutex.Lock()
	defer connMutex.Unlock()

	switch cs {
	case http.StateNew:
		log.Infof("CONN Opened: 0x%x\n", connPtr)
		conns[connPtr] = c

	case http.StateClosed:
		log.Infof("CONN Closed: 0x%x\n", connPtr)
		delete(conns, connPtr)
	}
}

///////////////////////////////////////////////////////////////////////////////////

type MultiProtoHttpCodec struct {
	conn *TcpConnWrapper
}

func NewMultiProtoHttpCodec(conn *TcpConnWrapper) *MultiProtoHttpCodec {
	return &MultiProtoHttpCodec{conn}
}

func (m *MultiProtoHttpCodec) Receive() (interface{}, error) {
	msgData, ok := <-m.conn.RecvChan
	log.Infof("%v,  ok: %v", msgData, ok)
	// fmt.Errorf()
	if !ok {
		return nil, fmt.Errorf("chan closed")
	}
	return msgData, nil
}

func (m *MultiProtoHttpCodec) Send(msg interface{}) error {
	m.conn.SendChan <- msg
	return nil
}

func (m *MultiProtoHttpCodec) Close() error {
	return m.conn.Close()
}

func (c *MultiProtoHttpCodec) Context() interface{} {
	return nil
}
