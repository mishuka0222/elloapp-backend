package net2

import (
	"sync"
)

const connectionMapNum = 32

type ConnectionManager struct {
	connectionMaps [connectionMapNum]connectionMap
	disposeFlag    bool
	disposeOnce    sync.Once
	disposeWait    sync.WaitGroup
}

type connectionMap struct {
	sync.RWMutex
	conns    map[uint64]Connection
	disposed bool
}

func NewConnectionManager() *ConnectionManager {
	manager := &ConnectionManager{}
	for i := 0; i < len(manager.connectionMaps); i++ {
		manager.connectionMaps[i].conns = make(map[uint64]Connection)
	}
	return manager
}

func (manager *ConnectionManager) Dispose() {
	manager.disposeOnce.Do(func() {
		manager.disposeFlag = true
		for i := 0; i < connectionMapNum; i++ {
			conns := &manager.connectionMaps[i]
			conns.Lock()
			conns.disposed = true
			for _, conn := range conns.conns {
				conn.Close()
			}
			conns.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *ConnectionManager) GetConnection(connID uint64) Connection {
	conns := &manager.connectionMaps[connID%connectionMapNum]
	conns.RLock()
	defer conns.RUnlock()

	conn, _ := conns.conns[connID]
	return conn
}

func (manager *ConnectionManager) putConnection(conn Connection) {
	conns := &manager.connectionMaps[conn.GetConnID()%connectionMapNum]

	conns.Lock()
	defer conns.Unlock()

	if conns.disposed {
		conn.Close()
		return
	}

	conns.conns[conn.GetConnID()] = conn
	manager.disposeWait.Add(1)
}

func (manager *ConnectionManager) delConnection(conn Connection) {
	if manager.disposeFlag {
		manager.disposeWait.Done()
		return
	}

	conns := &manager.connectionMaps[conn.GetConnID()%connectionMapNum]

	conns.Lock()
	defer conns.Unlock()

	delete(conns.conns, conn.GetConnID())
	manager.disposeWait.Done()
}
