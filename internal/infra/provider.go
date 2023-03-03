package infra

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var pool *sql.DB

func InitDB(driverName, dsn string) {
	var err error
	pool, err = sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	pool.SetMaxOpenConns(100)
	pool.SetMaxIdleConns(20)
	pool.SetConnMaxLifetime(100 * time.Second)

}
func GetDB() *sql.DB {
	return pool
}
