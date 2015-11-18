package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hulucat/conf"
)

func OpenMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", conf.Get("mysql"))
	if db != nil {
		db.SetMaxIdleConns(90)
		db.SetMaxOpenConns(90)
	}
	return db, err
}
