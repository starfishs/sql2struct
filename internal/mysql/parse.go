package mysqlparser

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenerateFile(ddl string) error {
	c, _ := ParseMysqlDDL(ddl)
	tmpDir := os.TempDir()
	fileName := filepath.Join(tmpDir, strings.ToLower(c.Name)+".go")
	fd, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer fd.Close()
	fd.Write([]byte(c.GenerateCode()))

	out := bytes.NewBuffer(nil)
	command := exec.Command("gofmt", fileName)
	command.Stdout = out

	err = command.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	os.MkdirAll("model", 0755)
	f, _ := os.Create("model/" + strings.ToLower(c.Name) + ".go")
	defer f.Close()

	_, err = f.Write([]byte(out.String()))
	if err != nil {
		fmt.Println("write ", err.Error())
	}
	return nil
}
