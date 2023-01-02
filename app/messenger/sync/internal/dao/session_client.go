package dao

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	session_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

// SessionOptions comet options.
type SessionOptions struct {
	RoutineSize uint64
	RoutineChan uint64
}

// Session is a gateway.
type Session struct {
	serverId       string
	client         session_client.SessionClient
	sessionChan    []chan interface{}
	sessionChanNum uint64
	options        SessionOptions
	ctx            context.Context
	cancel         context.CancelFunc
}

// process
func (c *Session) process(sessionChan chan interface{}) {
	var err error
	for {
		select {
		case updates, ok := <-sessionChan:
			if !ok {
				logx.Errorf("process error")
				return
			}

			switch r := updates.(type) {
			case *session.TLSessionPushSessionUpdatesData:
				_, err = c.client.SessionPushSessionUpdatesData(context.Background(), r)
				if err != nil {
					logx.Errorf("c.client.PushSessionUpdates(%s, %v, reply) serverId:%d error(%v)", r, c.serverId, err)
				}
			case *session.TLSessionPushUpdatesData:
				_, err = c.client.SessionPushUpdatesData(context.Background(), r)
				if err != nil {
					logx.Errorf("c.client.PushUpdates(%s, %v, reply) serverId:%d error(%v)", r, c.serverId, err)
				}
			case *session.TLSessionPushRpcResultData:
				_, err = c.client.SessionPushRpcResultData(context.Background(), r)
				if err != nil {
					logx.Errorf("c.client.PushRpcResult(%s, %v, reply) serverId:%d error(%v)", r, c.serverId, err)
				}
			default:
				logx.Errorf("invalid type: %#v", r)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Session) Close() (err error) {
	finish := make(chan bool)
	go func() {
		for {
			n := len(c.sessionChan)
			for _, ch := range c.sessionChan {
				n += len(ch)
			}
			if n == 0 {
				finish <- true
				return
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-finish:
		logx.Info("close session client finish")
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("close session(server:%s push:%d) timeout", c.serverId, len(c.sessionChan))
	}
	c.cancel()
	return
}

func (c *Session) PushUpdates(ctx context.Context, msg *session.TLSessionPushUpdatesData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

func (c *Session) PushSessionUpdates(ctx context.Context, msg *session.TLSessionPushSessionUpdatesData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

func (c *Session) PushRpcResult(ctx context.Context, msg *session.TLSessionPushRpcResultData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

// NewSession new a comet.
func NewSession(c zrpc.RpcClientConf, options SessionOptions) (*Session, error) {
	sess := &Session{
		serverId:    c.Endpoints[0],
		sessionChan: make([]chan interface{}, options.RoutineSize),
		options:     options,
	}

	cli, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("watchComet NewClient(%+v) error(%v)", c, err)
		return nil, err
	}
	sess.client = session_client.NewSessionClient(cli)
	sess.ctx, sess.cancel = context.WithCancel(context.Background())

	for i := uint64(0); i < options.RoutineSize; i++ {
		sess.sessionChan[i] = make(chan interface{}, options.RoutineChan)
		go sess.process(sess.sessionChan[i])
	}
	return sess, nil
}

func (d *Dao) watch(c zrpc.RpcClientConf) {
	sub, _ := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	update := func() {
		values := sub.Values()
		if len(values) == 0 {
			return
		}

		sessions := map[string]*Session{}
		for _, v := range values {
			if old, ok := d.sessionServers[v]; ok {
				sessions[v] = old
				continue
			}
			c.Endpoints = []string{v}
			// cli, err := zrpc.NewClient(c)
			cli, err := NewSession(c, SessionOptions{
				RoutineSize: d.conf.Routine.Size,
				RoutineChan: d.conf.Routine.Chan,
			})
			if err != nil {
				logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
				return
			}
			sessions[v] = cli
		}

		for key, old := range d.sessionServers {
			if _, ok := sessions[key]; !ok {
				old.cancel()
				logx.Infof("watchComet DelComet:%s", key)
			}
		}

		d.sessionServers = sessions
	}

	sub.AddListener(update)
	update()
}

func (d *Dao) PushUpdatesToSession(ctx context.Context, serverId string, msg *session.TLSessionPushUpdatesData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		// log.Info("push updates to serverId: (%s, %s)", serverId, msg.DebugString())
		return c.PushUpdates(ctx, msg)
	} else {
		logx.Errorf("not found k: %s, %v", serverId, d.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
	return
}

func (d *Dao) PushSessionUpdatesToSession(ctx context.Context, serverId string, msg *session.TLSessionPushSessionUpdatesData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		// logx.Info("push session updates to serverId: (%s, %s)", serverId, msg.DebugString())
		return c.PushSessionUpdates(ctx, msg)
	} else {
		logx.Errorf("not found k: %s, %v", serverId, d.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
	return
}

func (d *Dao) PushRpcResultToSession(ctx context.Context, serverId string, msg *session.TLSessionPushRpcResultData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		// log.Debugf("push rpc result to serverId: (%s, %s)", serverId, msg.DebugString())
		return c.PushRpcResult(ctx, msg)
	} else {
		logx.Errorf("not found k: %s, %v", serverId, d.sessionServers)
		return fmt.Errorf("not found k: %s", serverId)
	}
	return
}
