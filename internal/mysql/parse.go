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
	//ddl = "CREATE TABLE `project` (\n  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL COMMENT '项目名称',\n  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1:未上线, 2:已上线,3:上线失败',\n  `valid_from` datetime DEFAULT NULL COMMENT '项目上线时间',\n  `failed_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '上线失败原因',\n  `top_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '置顶标记 :置顶',\n  `allow_empty_payment_setting` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否允许无付款设置 1: 不允许 2:允许',\n  `allow_empty_collection_setting` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否允许无收款设置 1: 不允许 2:允许',\n  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',\n  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',\n  `on_top_time` datetime(6) DEFAULT NULL COMMENT '项目置顶时间',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `name` (`name`)\n) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8mb4 COMMENT='项目表'"
	c, _ := ParseMysqlDDL(ddl)
	fmt.Println(c)
	//fmt.Println(c.GenerateCode())
	tmpDir := os.TempDir()
	fmt.Println(tmpDir)
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
	f, _ := os.Create("model/" + strings.ToLower(c.Name) + ".go")
	defer f.Close()
	fmt.Println(out.String())
	_, err = f.Write([]byte(out.String()))
	if err != nil {
		fmt.Println("write ", err.Error())
	}
	return nil
}
