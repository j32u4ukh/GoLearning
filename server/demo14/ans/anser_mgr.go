package ans

import (
	"net"
	"sync"
)

type AnserManager struct {
	AnserMap map[string]*Anser
	amMux    sync.Mutex
	ip       net.IP
}

func (am *AnserManager) Init(addr string) {
	am.ip = net.ParseIP(addr)
}
