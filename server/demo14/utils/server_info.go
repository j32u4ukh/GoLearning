package utils

import (
	"fmt"
	"net"
)

var ServerAddress map[string]ServerInfo

type ServerInfo struct {
	Ip   string
	Port int
	// TODO: 之後直接 GS 對應 127.0.0.1:8080 所對應的 *net.TCPAddr 物件
	TcpAddr *net.TCPAddr
}

func init() {
	ServerAddress = make(map[string]ServerInfo)
	ServerAddress["GS"] = ServerInfo{Ip: "127.0.0.1", Port: 8080}
}

func GetServerInfo(s string) ServerInfo {
	if si, ok := ServerAddress[s]; ok {
		return si
	}
	return ServerInfo{}
}

func (s ServerInfo) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
