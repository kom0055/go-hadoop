package ipc

import "sync"

var (
	connectionPoolWrapper = connectionPool{connections: map[connIdentity]*connection{}}
)

type (
	connectionPool struct {
		mu          sync.RWMutex
		connections map[connIdentity]*connection
	}
)

func (p *connectionPool) GetConnection(connectionId *connIdentity) (*connection, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	con, ok := connectionPoolWrapper.connections[*connectionId]
	return con, ok
}

func (p *connectionPool) SetConnection(connectionId *connIdentity, conn *connection) {
	p.mu.Lock()
	defer p.mu.Unlock()
	connectionPoolWrapper.connections[*connectionId] = conn
}
