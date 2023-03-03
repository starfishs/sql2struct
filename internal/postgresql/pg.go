package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gangming/sql2struct/config"
	"github.com/gangming/sql2struct/internal/infra"
	"github.com/gangming/sql2struct/internal/table"
	"github.com/gangming/sql2struct/utils"
)

type pgParser struct {
}
type PgGenerator struct{}

func NewPgParser() *PgGenerator {
	return &PgGenerator{}
}

const (
	fieldInfoSql = `select
	c.relname as table_name,
	a.attname as field_name,
-- 	(case
-- 		when a.attnotnull = true then true
-- 		else false end) as 'isnotnull',
	(case
		when (
		select
			count(pg_constraint.*)
		from
			pg_constraint
		inner join pg_class on
			pg_constraint.conrelid = pg_class.oid
		inner join pg_attribute on
			pg_attribute.attrelid = pg_class.oid
			and pg_attribute.attnum = any(pg_constraint.conkey)
		inner join pg_type on
			pg_type.oid = pg_attribute.atttypid
		where
			pg_class.relname = c.relname
			and pg_constraint.contype = 'p'
			and pg_attribute.attname = a.attname) > 0 then true
		else false end) as is_primary_key,
	concat_ws('', t.typname) as field_type,
-- 	(case
-- 		when a.attlen > 0 then a.attlen
-- 		when t.typname='bit' then a.atttypmod
-- 		else a.atttypmod - 4 end) as field_length,

-- 	 col.is_identity	as is_auto_increment,

	 coalesce(col.column_default,'')	as default_value,
	coalesce(
	    (select description from pg_description where objoid = a.attrelid
	and objsubid = a.attnum),'') as field_comment
from
	pg_class c,
	pg_attribute a ,
	pg_type t,
	information_schema.columns as col
where
	c.relname = $1
	and a.attnum>0
	and a.attrelid = c.oid
	and a.atttypid = t.oid
	and col.table_name=c.relname and col.column_name=a.attname
order by
	c.relname desc,
	a.attnum asc`
)

var PgType2GoType = map[string]string{
	"smallint":          "int16",
	"integer":           "int32",
	"bigint":            "int64",
	"int8":              "int64",
	"int2":              "int16",
	"int4":              "int32",
	"real":              "float32",
	"double precision":  "float64",
	"boolean":           "bool",
	"bool":              "bool",
	"bytea":             "[]byte",
	"varchar":           "string",
	"json":              "string",
	"character varying": "string",
	"character":         "string",
	"text":              "string",
	"uuid":              "string",
	"inet":              "string",
	"macaddr":           "string",
	"timestamp":         "time.Time",
	"timestamptz":       "time.Time",
	"date":              "time.Time",
	"interval":          "time.Duration",
	"jsonb":             "map[string]interface{}",
}

func (p *PgGenerator) Run() error {
	parser := &pgParser{}
	tables, err := parser.GetAllTables()
	if err != nil {
		return err
	}
	for _, t := range tables {
		err := parser.ParseFields(t)
		if err != nil {
			return err
		}
		err = t.GenerateFile()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *pgParser) GetAllTables() ([]*table.Table, error) {
	var (
		rows   *sql.Rows
		err    error
		tables []*table.Table
	)

	if len(config.Cnf.Tables) > 0 {
		tablesStr := strings.Join(config.Cnf.Tables, "','")
		rows, err = infra.GetDB().Query("select relname as tabname,COALESCE(cast(obj_description(relfilenode,'pg_class') as varchar),'') as comment from pg_class c where relname in ($1)", tablesStr)
	} else {
		rows, err = infra.GetDB().Query("select relname as tabname,COALESCE(cast(obj_description(relfilenode,'pg_class') as varchar),'') as comment from pg_class c where relname in (select tablename from pg_tables where schemaname='public' and position('_2' in tablename)=0)")
	}
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			tableName string
			comment   string
		)

		err = rows.Scan(&tableName, &comment)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, &table.Table{
			Package:            config.Cnf.PackageName,
			Name:               tableName,
			UpperCamelCaseName: utils.Underline2UpperCamelCase(tableName),
			Comment:            comment,
		})
		fmt.Println(tableName)
	}
	return tables, err
}
func (p *pgParser) ParseFields(t *table.Table) error {
	rows, err := infra.GetDB().Query(fieldInfoSql, t.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			isPrimaryKey bool
			tableName    string
			fieldName    string
			fieldType    string
			defaultValue string
			fieldComment string
		)

		err = rows.Scan(&tableName, &fieldName, &isPrimaryKey, &fieldType, &defaultValue, &fieldComment)
		fieldComment = fieldType
		if err != nil {
			log.Fatal(err)
		}
		fType, ok := PgType2GoType[fieldType]
		if !ok {
			fType = "interface{}"
		}
		if strings.Contains(fType, "time") {
			t.ContainsTimeField = true
		}
		if tableName == t.Name {
			f := table.Field{
				IsPK:               isPrimaryKey,
				Name:               fieldName,
				UpperCamelCaseName: utils.Underline2UpperCamelCase(fieldName),
				Type:               fType,
				Comment:            fieldComment,
			}
			if config.Cnf.WithDefaultValue {
				f.DefaultValue = defaultValue
			}
			t.Fields = append(t.Fields, f)
		}
	}
	return nil
}
