package main

import (
	"flag"
	"fmt"
	"github.com/gangming/sql2struct/internal/infra"
	mysqlparser "github.com/gangming/sql2struct/internal/mysql"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "dsn", "", "eg: -dsn=root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	flag.Parse()
	fmt.Println(dsn)
	if dsn == "" {
		panic("dsn is required eg: -dsn=root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	}
	infra.InitDBMysql(dsn)

	ddls, err := GetDDLs()
	if err != nil {
		panic(err)
	}
	for _, ddl := range ddls {
		mysqlparser.GenerateFile(ddl)
	}
}

func GetDDLs() ([]string, error) {
	var result []string
	tables := GetTables()
	for _, tableName := range tables {
		rows, err := infra.GetDB().Query("show create table " + tableName)
		if err != nil {
			panic(err)
		}

		if rows.Next() {
			var r string
			err := rows.Scan(&tableName, &r)
			if err != nil {
				panic(err)
			}
			result = append(result, r)
		}
	}

	return result, nil
}
func GetTables() []string {
	var result []string
	rows, err := infra.GetDB().Query("show tables")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			panic(err)
		}
		result = append(result, r)
	}
	return result
}
