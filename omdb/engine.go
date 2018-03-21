package omdb

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"sync"
)

type ConfEngine struct {
	DBEngine   *xorm.Engine
	gConfFiles []string
}

var g_conf_engine *ConfEngine
var g_mutex sync.Mutex

func (eng *ConfEngine) init() int {
	if dbeng, err := CreateDBEngine(); err == nil {
		eng.DBEngine = dbeng
		eng.init_tables()
		return 0
	}
	fmt.Println("[ERR] Init() failure!")
	return -1
}

func (eng *ConfEngine) init_tables() {
	return
}

func GetDB() *xorm.Engine {
	g_mutex.Lock()
	defer g_mutex.Unlock()

	if g_conf_engine == nil {
		g_conf_engine = &ConfEngine{}
	}

	g_conf_engine.init()

	return g_conf_engine.DBEngine
}
