package infra

import (
	"database/sql"
	"time"
)

var pool *sql.DB

func InitDBMysql() {
	var err error
	pool, err = sql.Open("mysql", "planet_finance:finance123456@tcp(sandbox-common-finance-mysql.cluster-choxzj9zxm2u.rds.cn-northwest-1.amazonaws.com.cn:3306)/settle_setting?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		panic(err)
	}
	pool.SetMaxOpenConns(100)
	pool.SetMaxIdleConns(20)
	pool.SetConnMaxLifetime(100 * time.Second)

}
