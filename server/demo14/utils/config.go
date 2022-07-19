package utils

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
)

var instance *Config
var once sync.Once

type Config struct {
	SendCode SendCode `yaml:"SendCode"`
	Addr     Addr     `yaml:"Addr"`
	Conn     Conn     `yaml:"Conn"`
}

type SendCode struct {
	Close        byte `yaml:"Close"`
	RegisterReq  byte `yaml:"RegisterReq"`
	RegisterRes  byte `yaml:"RegisterRes"`
	HeartBeatReq byte `yaml:"HeartBeatReq"`
	HeartBeatRes byte `yaml:"HeartBeatRes"`
	Req          byte `yaml:"Req"`
	Res          byte `yaml:"Res"`
	ProtobufReq  byte `yaml:"ProtobufReq"`
	ProtobufRes  byte `yaml:"ProtobufRes"`
}

type Addr struct {
	GS string `yaml:"GS"`
	AS string `yaml:"AS"`
}

type Conn struct {
	RetryTimes byte    `yaml:"RetryTimes"`
	Alternal   float32 `yaml:"Alternal"`
}

func (c *Config) Init() error {
	var b []byte
	var err error
	b, err = ioutil.ReadFile("data/config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &instance)
	return err
}

func GetConfig() *Config {
	if instance == nil {
		once.Do(func() {
			instance = &Config{}
			err := instance.Init()
			if err != nil {
				fmt.Println("Loading config failed, err:", err.Error())
			}
		})
	}

	return instance
}

func GetSendCode() SendCode {
	return GetConfig().SendCode
}

func GetAddr() Addr {
	return GetConfig().Addr
}

func GetConn() Conn {
	return GetConfig().Conn
}
