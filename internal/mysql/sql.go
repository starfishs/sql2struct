package mysqlparser

import (
	"fmt"
	"strings"

	"github.com/gangming/sql2struct/config"
	"github.com/gangming/sql2struct/internal/infra"
	"github.com/gangming/sql2struct/internal/table"
	"github.com/gangming/sql2struct/utils"
)

var MysqlType2GoType = map[string]string{
	"int":       "int64",
	"tinyint":   "uint8",
	"decimal":   "float64",
	"bigint":    "int64",
	"varchar":   "string",
	"char":      "string",
	"text":      "string",
	"date":      "time.Time",
	"time":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"json":      "string",
}

type mysqlParser struct {
}
type MysqlGenerator struct {
}

func NewMysqlGenerator() *MysqlGenerator {
	return &MysqlGenerator{}
}

func (m *mysqlParser) ParseMysqlDDL(s string) (table.Table, error) {
	lines := strings.Split(s, "\n")
	var t table.Table
	t.Package = config.Cnf.PackageName
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "CREATE TABLE") {
			tableName := strings.Split(line, "`")[1]
			t.Name = config.Cnf.TablePrefix + tableName
			t.UpperCamelCaseName = utils.Underline2UpperCamelCase(t.Name)
			continue
		}
		if strings.Contains(line, "ENGINE") && strings.Contains(line, "COMMENT=") {
			t.Comment = strings.Trim(strings.Split(line, "COMMENT='")[1], "'")
			fmt.Println(t.Comment)
			continue
		}
		if line[0] == '`' {
			field := table.Field{}
			field.Name = strings.Split(line, "`")[1]
			field.UpperCamelCaseName = utils.Underline2UpperCamelCase(field.Name)
			field.Type = strings.TrimRightFunc(strings.Split(line, " ")[1], func(r rune) bool {
				return r < 'a' || r > 'z'
			})
			field.Type = MysqlType2GoType[field.Type]
			if strings.Contains(field.Type, "time") {
				t.ContainsTimeField = true
			}
			if strings.Contains(line, "COMMENT") {
				field.Comment = strings.Trim(strings.Split(line, "COMMENT '")[1], "',")
			}
			if strings.Contains(line, "DEFAULT'") {
				field.DefaultValue = strings.Split(line, "DEFAULT ")[1]
			}
			if strings.Contains(line, "PRIMARY KEY") {
				field.IsPK = true
			}

			t.Fields = append(t.Fields, field)

		}

	}
	return t, nil
}

func (m *mysqlParser) GetDDLs() ([]string, error) {
	var result []string
	tables := m.GetTables()
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
func (m *mysqlParser) GetTables() []string {
	if len(config.Cnf.Tables) > 0 {
		return config.Cnf.Tables
	}
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
func (m *MysqlGenerator) Run() error {
	p := &mysqlParser{}
	ddls, err := p.GetDDLs()
	if err != nil {
		return err
	}
	for _, ddl := range ddls {
		c, err := p.ParseMysqlDDL(ddl)
		if err != nil {
			return err
		}
		err = c.GenerateFile()
		if err != nil {
			return err
		}
	}
	return nil
}
