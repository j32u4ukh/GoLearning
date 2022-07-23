package utils

import (
	"fmt"
	"net"
)

var ServerAddress map[string]*ServerInfo

type ServerInfo struct {
	ip   string
	port int
	addr string
	// TODO: 之後直接 GS 對應 127.0.0.1:8080 所對應的 *net.TCPAddr 物件
	tcpAddr *net.TCPAddr
}

func init() {
	ServerAddress = make(map[string]*ServerInfo)
	ServerAddress["GS"] = newServerAddress("127.0.0.1", 8080)
}

func newServerAddress(ip string, port int) *ServerInfo {
	s := &ServerInfo{ip: ip, port: port}
	s.addr = fmt.Sprintf("%s:%d", ip, port)
	return s
}

func GetServerInfo(s string) *ServerInfo {
	if si, ok := ServerAddress[s]; ok {
		return si
	}
	return &ServerInfo{}
}

func (s ServerInfo) GetAddress() string {
	return s.addr
}
