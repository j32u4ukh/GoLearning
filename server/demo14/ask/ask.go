package ask

// golang實現帶有心跳檢測的tcp長連接
// server

import (
	"GoLearning/server/demo14/ask/site"
	"GoLearning/server/demo14/utils"
	"fmt"
	"sync"
)

// Call to registered server
func call() {
	wg := &sync.WaitGroup{}
	wg.Add(len(GetAskerManager().Connects))

	for _, c := range GetAskerManager().Connects {
		go GetAskerManager().Connect(c, wg)
	}

	wg.Wait()
}

func Init() {
	GetAskerManager().Connects = append(GetAskerManager().Connects, utils.GetAddr().GS)
	call()
	tsg := site.TGS{}
	tsg.Init()
	GetAskerManager().RegisterFunc(utils.GetAddr().GS, "Message", tsg.Callback)
}

func main() {
	GetAskerManager().Send(utils.GetAddr().GS, []byte{utils.GetSendCode().Req, 'F', 'i', 'r', 's', 't'})
	GetAskerManager().Send(utils.GetAddr().GS, []byte{utils.GetSendCode().Req, 'S', 'e', 'c', 'o', 'n', 'd'})

	end := <-GetAskerManager().Cch
	fmt.Println("End of connection.", end)
}
