package utils

import "fmt"

var ServerAddress map[string]ServerInfo

type ServerInfo struct {
	Ip   string
	Port int
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

func init() {
	ServerAddress = make(map[string]ServerInfo)
	ServerAddress["GS"] = ServerInfo{Ip: "127.0.0.1", Port: 8080}
}
