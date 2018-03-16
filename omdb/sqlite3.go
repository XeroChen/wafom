package omdbsqlite

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//"github.com/mattn/go-sqlite3"
)

func CreateDBEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "./webapp_sqlite3.db")
	if err == nil {
		engine.SetMapper(core.SameMapper{})
		tbMapper := core.NewPrefixMapper(core.SameMapper{}, "waf_")
		engine.SetTableMapper(tbMapper)
		return engine, nil
	}
	fmt.Println("xorm.NewEngine failed with: ", err.Error())
	return nil, err
}
