package ans

import (
	"GoLearning/server/demo16/config"
	"GoLearning/server/demo16/log16"
)

func Init() {
	log16.Logger().Debug("Ans.Init()")
	GetAnserManager().listenTCP(config.GetAddr().GS)
	log16.Logger().Debug("After Ans.Init()")
}
