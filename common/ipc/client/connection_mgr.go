package ipc

import (
	"context"
	"net"
	"sync"

	"github.com/kom0055/go-hadoop/common/defined"
	"github.com/kom0055/go-hadoop/common/log"
	uuid "github.com/nu7hatch/gouuid"
)

var (
	connectionMgr = connectionManager{connections: map[connIdentity]*connection{}}
)

type (
	connectionManager struct {
		mu          sync.RWMutex
		connections map[connIdentity]*connection
	}
)

func (p *connectionManager) GetConnection(ctx context.Context, clientId *uuid.UUID, connectionId connIdentity) (*connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	conn, ok := p.connections[connectionId]
	if ok {
		return conn, nil
	}

	//addr, _ := net.ResolveTCPAddr("tcp", c.ServerAddress)
	//tcpConn, err := net.DialTCP("tcp", nil, addr)
	tcpConn, err := (&net.Dialer{}).DialContext(ctx, "tcp", connectionId.address)
	if err != nil {
		log.Warnf("error: %v", err)
		return nil, err
	}
	log.Infof("Successfully connected: %v ", connectionId)

	// TODO: Ping thread

	// Set tcp no-delay
	//if tc, ok := tcpConn.(*net.TCPConn); ok {
	//	_ = tc.SetNoDelay(tcpNoDelay)
	//}
	conn = &connection{tcpConn}
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	authProtocol := defined.AUTH_PROTOCOL_NONE
	if _, found := findUsableTokenForService(connectionId.address); found {
		log.Infof("found token for service: %s", connectionId.address)
		authProtocol = defined.AUTH_PROTOCOL_SASL
	}

	if err = conn.writeConnectionHeader(authProtocol); err != nil {
		log.Infof("writeConnectionHeader failed: %v", err)
		return nil, err
	}

	if authProtocol == defined.AUTH_PROTOCOL_SASL {
		log.Infof("attempting SASL negotiation.")

		if err = conn.negotiateSimpleTokenAuth(connectionId.address); err != nil {
			log.Warnf("failed to complete SASL negotiation!: %v", err)
			return nil, err
		}

	} else {
		log.Infof("no usable tokens. proceeding without auth.")
	}

	if err = conn.writeConnectionContext(clientId, connectionId, authProtocol); err != nil {
		log.Infof("writeConnectionContext failed: %v", err)
		return nil, err
	}

	p.connections[connectionId] = conn

	return conn, nil

}

func (p *connectionManager) RmvConnection(connectionId connIdentity) {
	p.mu.Lock()
	defer p.mu.Unlock()
	conn, exist := p.connections[connectionId]
	if !exist {
		return
	}
	defer conn.Close()
	delete(p.connections, connectionId)

}
