package ask

// golang實現帶有心跳檢測的tcp長連接
// server

import (
	"GoLearning/server/demo16/ask/site"
	"GoLearning/server/demo16/config"
	"sync"
)

// func main() {
// 	GetAskerManager().Send(config.GetAddr().GS, []byte{config.GetSendCode().Req, 'F', 'i', 'r', 's', 't'})
// 	GetAskerManager().Send(config.GetAddr().GS, []byte{config.GetSendCode().Req, 'S', 'e', 'c', 'o', 'n', 'd'})

// 	end := <-GetAskerManager().Cch
// 	fmt.Println("End of connection.", end)
// }

func Init() {
	GetAskerManager().Connects = append(GetAskerManager().Connects, config.GetAddr().GS)
	call()
	tsg := site.TGS{}
	tsg.Init()
	GetAskerManager().RegisterFunc(config.GetAddr().GS, "Message", tsg.Callback)
}

// Call to registered server
func call() {
	wg := &sync.WaitGroup{}
	wg.Add(len(GetAskerManager().Connects))

	for _, c := range GetAskerManager().Connects {
		go GetAskerManager().Connect(c, wg)
	}

	wg.Wait()
}
