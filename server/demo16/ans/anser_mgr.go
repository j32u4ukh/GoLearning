package ans

import (
	"GoLearning/server/demo16/log16"
	"GoLearning/server/demo16/utils"
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
	log16.Logger().Debug(fmt.Sprintf("Listen to %s, addr: %s\n", s, addr))
	var ask *Anser
	var ok bool

	if ask, ok = am.anserMap[addr]; !ok {
		ip := si.GetIp()
		port := si.GetPort()
		ask = am.addAnswer(addr, ip, port)
	}

	err := ask.ListenTCP()

	if err != nil {
		log16.Logger().Error(err.Error())
		return err

	} else {
		log16.Logger().Debug("ask.Run()")
		go ask.Run()
	}

	log16.Logger().Debug("Waitting for ask.stopCh")
	<-ask.stopCh
	log16.Logger().Debug("Got ask.stopCh")

	return err
}

func (am *AnserManager) run() {

}

func (am *AnserManager) addAnswer(addr string, ip string, port int) *Anser {
	am.amMux.Lock()
	netIp := net.ParseIP(ip)
	am.anserMap[addr] = NewAnser(netIp, port)
	am.amMux.Unlock()
	return am.anserMap[addr]
}
