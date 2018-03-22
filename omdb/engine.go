package omdb

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"sync"
)

var gEngine *xorm.Engine
var gMutex sync.Mutex

func initDB() *xorm.Engine {
	dbeng, err := CreateDBEngine()
	if err != nil {
		fmt.Println("[ERR] Init() failure! ", err.Error())
		return nil
	}
	return dbeng
}

func GetDB() *xorm.Engine {
	gMutex.Lock()
	defer gMutex.Unlock()

	if gEngine == nil {
		gEngine = initDB()
	}
	return gEngine
}

func CloseDB(*xorm.Engine) {
	gMutex.Lock()
	defer gMutex.Unlock()

	if gEngine != nil {
		gEngine.Close()
		gEngine = nil
	}
}
