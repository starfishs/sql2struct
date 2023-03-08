package driver

import (
	mysqlparser "github.com/gangming/sql2struct/internal/mysql"
	"github.com/gangming/sql2struct/internal/postgresql"
	"github.com/gangming/sql2struct/utils"
)

type ModelGenerator interface {
	Run() error
}

func NewSqlDriverGenerator(driverName string) ModelGenerator {

	if driverName == "mysql" {
		return mysqlparser.NewMysqlGenerator()
	}
	if driverName == "postgres" {
		return postgresql.NewPgParser()
	}
	utils.PrintRedf("unsupported driver %s", driverName)
	return nil
}
