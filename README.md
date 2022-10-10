[![Build Status](https://travis-ci.org/gangming/sql2struct.svg?branch=main)](https://travis-ci.org/gangming/sql2struct)
[![Go Report Card](https://goreportcard.com/badge/github.com/gangming/sql2struct)](https://goreportcard.com/report/github.com/gangming/sql2struct)
[![GoDoc](https://godoc.org/github.com/gangming/sql2struct?status.svg)](https://godoc.org/github.com/gangming/sql2struct)
[![codecov](https://codecov.io/gh/gangming/sql2struct/branch/main/graph/badge.svg)](https://codecov.io/gh/gangming/sql2struct)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
# sql2struct
mysql database to golang struct for gorm model

# install
```shell
go install github.com/gangming/sql2struct@latest
```



# usage
```shell
sql2struct -dsn="root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
```

