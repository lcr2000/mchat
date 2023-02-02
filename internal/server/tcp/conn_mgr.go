package tcp

import (
	"sync"
)

var connMgr *ConnMgr

func init() {
	connMgr = &ConnMgr{connMgr: make(map[string]*Conn)}
}

func GetConnMgr() *ConnMgr {
	return connMgr
}

type ConnMgr struct {
	connMgr map[string]*Conn
	sync.Mutex
}

func (cm *ConnMgr) Add(id string, conn *Conn) {
	cm.Lock()
	defer cm.Unlock()
	cm.connMgr[id] = conn
}

func (cm *ConnMgr) Remove(id string) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.connMgr, id)
}

func (cm *ConnMgr) Get(id string) (*Conn, bool) {
	cm.Lock()
	defer cm.Unlock()
	conn, ok := cm.connMgr[id]
	return conn, ok
}

func (cm *ConnMgr) GetAll() []*Conn {
	cm.Lock()
	defer cm.Unlock()
	conns := make([]*Conn, 0, len(cm.connMgr))
	for _, conn := range cm.connMgr {
		conns = append(conns, conn)
	}
	return conns
}

func (cm *ConnMgr) Close() {
	cm.Lock()
	defer cm.Unlock()
	cm.connMgr = make(map[string]*Conn)
}
