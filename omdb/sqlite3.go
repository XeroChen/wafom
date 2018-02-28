package omdbsqlite

import (
	"github.com/go-xorm/xorm"
	"github.com/mattn/go-sqlite3"
)

type DBAttr struct {
	dbi interface{}
}

type DBMethod interface {
	Conn(fd interface{})
	Exec(sqlstmt string)
	DisConn()
}

type DBObj struct {
	DBAttr
	DBMethod
}

func CreateSQLite3Engine() *xorm.Engine {
	var err error
	var engine *xorm.Engine
	engine, err = xorm.NewEngine("sqlite3", "./webapp_sqlite3.db")
	if engine != nil {
		return engine
	}
	return nil
}
