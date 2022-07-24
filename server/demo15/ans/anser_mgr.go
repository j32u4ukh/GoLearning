package ans

import (
	"GoLearning/server/demo15/log15"
	"GoLearning/server/demo15/utils"
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
	am.anserMap = make(map[string]*Anser)
}

func (am *AnserManager) listenTCP(s string) error {
	// fmt.Printf("Listen to %s\n", s)
	si := utils.GetServerInfo(s)
	addr := si.GetAddress()
	log15.Logger().Debug(fmt.Sprintf("Listen to %s, addr: %s\n", s, addr))
	var ask *Anser
	var ok bool

	if ask, ok = am.anserMap[addr]; !ok {
		ip := si.GetIp()
		port := si.GetPort()
		ask = am.addAnswer(addr, ip, port)
	}

	err := ask.ListenTCP()

	if err != nil {
		log15.Logger().Error(err.Error())
		return err

	} else {
		log15.Logger().Debug("ask.Run()")
		go ask.Run()
	}

	log15.Logger().Debug("Waitting for ask.stopCh")
	<-ask.stopCh
	log15.Logger().Debug("Got ask.stopCh")

	return err
}

func (am *AnserManager) addAnswer(addr string, ip string, port int) *Anser {
	am.amMux.Lock()
	netIp := net.ParseIP(ip)
	am.anserMap[addr] = NewAnser(netIp, port)
	am.amMux.Unlock()
	return am.anserMap[addr]
}
