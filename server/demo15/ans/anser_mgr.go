package ans

import (
	"GoLearning/server/demo14/utils"
	"fmt"
	"net"
	"sync"
)

var instance *AnserManager
var once sync.Once

type AnserManager struct {
	anserMap map[string]*Anser
	amMux    sync.Mutex
}

func GetAnserManager() *AnserManager {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating AnserManager Instance Now")
			instance = &AnserManager{}
			instance.Init()
		})
	}

	return instance
}

func (am *AnserManager) Init() {
	am.anserMap = map[string]*Anser{}
}

func (am *AnserManager) listenTCP(s string, port int) error {
	si := utils.GetServerInfo(s)
	addr := si.GetAddress()
	var ask *Anser
	var ok bool

	if ask, ok = am.anserMap[addr]; !ok {
		// ask = am.addAnswer(addr, ip, port)
	}

	err := ask.ListenTCP()

	// if err != nil {

	// } else {

	// }

	return err
}

func (am *AnserManager) addAnswer(addr string, ip string, port int) *Anser {
	am.amMux.Lock()
	netIp := net.ParseIP(ip)
	am.anserMap[addr] = NewAnser(netIp, port)
	am.amMux.Unlock()
	return am.anserMap[addr]
}
