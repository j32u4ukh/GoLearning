package ans

import (
	"GoLearning/server/demo15/config"
	"GoLearning/server/demo15/log15"
)

func Init() {
	log15.Logger().Debug("Ans.Init()")
	GetAnserManager().listenTCP(config.GetAddr().GS)
	log15.Logger().Debug("After Ans.Init()")
	// server := &Ans{}
	// err := server.Init("127.0.0.1", 8080)

	// if err != nil {
	// 	fmt.Println("監聽端口失敗:", err.Error())
	// 	return
	// }

	// fmt.Println("已初始化連接，等待客戶端連接...")

	// // listen to client
	// server.Run()
}
