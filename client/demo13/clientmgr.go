package main

import (
	"fmt"
	"sync"
)

var instance *Client13Manager

var once sync.Once

type Client13Manager struct {
	ClientMap map[string]*Client13
	cmMux     sync.Mutex

	cacheMap map[string][][]byte
	cacheMux sync.Mutex

	Dch chan bool
	Cch chan string
}

func GetClient13Manager() *Client13Manager {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating Single Instance Now")
			instance = &Client13Manager{}
			instance.Init()
		})
	}

	return instance
}

func (cc *Client13Manager) Init() {
	cc.ClientMap = make(map[string]*Client13)
	cc.Cch = make(chan string, 128)
}

func (cc *Client13Manager) Send(s string, msg []byte) {
	if c, ok := cc.ClientMap[s]; ok {
		c.send(msg)
	} else {
		var cache [][]byte
		var ok bool
		if cache, ok = cc.cacheMap[s]; !ok {
			cc.cacheMap[s] = make([][]byte, 0)
		}

		cc.cacheMap[s] = append(cache, msg)
	}
}

// func (cc *Client13Manager) connect(ip string, port int) {
// 	addr := fmt.Sprintf("%v:%d", ip, port)

// 	client := &Client13{Pch: &cc.Cch, Addr: addr}
// 	client.Init(ip, port)
// 	client.Run()
// 	fmt.Println("After client.Run(), Addr:", client.Addr)

// 	cc.Cch <- client.Addr
// 	fmt.Println("cc.Cch <- ", client.Addr)
// }

func (cc *Client13Manager) Connect(s string) {
	si := GetServerInfo(s)
	addr := si.GetAddress()
	if _, ok := cc.ClientMap[addr]; !ok {
		client := &Client13{Pch: &cc.Cch, Addr: addr}
		client.Init(si.Ip, si.Port)
		cc.addNewConnection(addr, client)

		client.Run()
		fmt.Println("After client.Run(), Addr:", client.Addr)

		if cc.Cch != nil {
			fmt.Println("cc.Cch != nil")
		} else {
			fmt.Println("cc.Cch is nil")
		}

		cc.Cch <- client.Addr
		fmt.Println("cc.Cch <- ", client.Addr)
		return
	}
}

func (cc *Client13Manager) addNewConnection(addr string, c *Client13) {
	cc.cmMux.Lock()
	cc.ClientMap[addr] = c
	cc.cmMux.Unlock()
}
