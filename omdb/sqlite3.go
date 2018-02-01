package omdbsqlite

import "github.com/mattn/go-sqlite3"

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
