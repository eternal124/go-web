package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@/test?parseTime=true")
	if err != nil {
		log.Fatal("database init error: ", err)
	}

	SqlDB.Ping()
}
