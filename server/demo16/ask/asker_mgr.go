package ask

import (
	"GoLearning/server/demo16/log16"
	"GoLearning/server/demo16/utils"
	"fmt"
	"sync"
)

var instance *AskerManager
var once sync.Once

type AskerManager struct {
	AskerMap map[string]*Asker
	amMux    sync.Mutex

	Connects []string

	Dch chan bool
	Cch chan string
}

func GetAskerManager() *AskerManager {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating AskerManager Instance Now")
			instance = &AskerManager{}
			instance.Init()
		})
	}

	return instance
}

func (am *AskerManager) Init() {
	am.AskerMap = make(map[string]*Asker)
	am.Cch = make(chan string, 128)
}

func (am *AskerManager) Run() {

}

func (am *AskerManager) Send(s string, msg []byte) {
	si := utils.GetServerInfo(s)
	addr := si.GetAddress()
	if c, ok := am.AskerMap[addr]; ok {
		c.send(msg)
	}
}

// TODO: 每次連線傳入再次嘗試次數，避免無限循環
func (am *AskerManager) Connect(s string, wg *sync.WaitGroup) {
	si := utils.GetServerInfo(s)
	addr := si.GetAddress()
	asker, ok := am.AskerMap[addr]

	if !ok {
		// 若連往 addr 的連線尚不存在

		// 建立 Asker 實體
		// asker = &Asker{Addr: addr}
		asker = NewAsker(addr)

		// 紀錄連線關係
		am.addNewConnection(addr, asker)

		// 初始化
		// asker.Init()
	}

	// 連線
	err := asker.Connect()

	if err != nil {
		// 連線失敗，再次連線
		fmt.Printf("連線至 %s 時發生錯誤\n", addr)
		log16.Logger().Error(fmt.Sprintf("連線至 %s 時發生錯誤, err: %+v\n", addr, err.Error()))
		am.Connect(s, wg)

	} else {
		// 維持運行
		asker.Run(wg)

		if asker.isShutdown {
			// 因中斷連線而到這裡，從 AskerMap 中移除連線資訊
			delete(am.AskerMap, addr)

			// 將中斷伺服器的資訊傳給上層物件
			am.Cch <- asker.Addr
		} else {
			// 異常斷線，需再次連線
			fmt.Printf("連線至 %s 異常斷線，需再次連線", addr)
			am.Connect(s, nil)
		}
	}
}

func (am *AskerManager) addNewConnection(addr string, c *Asker) {
	am.amMux.Lock()
	am.AskerMap[addr] = c
	am.amMux.Unlock()
}

func (am *AskerManager) RegisterFunc(s string, msg string, callback func([]byte)) {
	si := utils.GetServerInfo(s)
	addr := si.GetAddress()

	if c, ok := am.AskerMap[addr]; ok {
		c.RegisterFunc(msg, callback)
	} else {
		fmt.Printf("addr %s is not exist\n", addr)
	}
}
