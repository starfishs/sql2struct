package mysqlparser

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/gangming/sql2struct/config"
)

func TestMysqlGenerateFile(t *testing.T) {
	type args struct {
		ddl string
	}
	input, err := os.ReadFile("testdata/user.input")
	if err != nil {
		t.Fatal("read testdata/user.input failed", err)
	}
	golden, err := os.ReadFile("testdata/user.golden")
	if err != nil {
		t.Fatal("read testdata/user.golden failed", err)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "testdata",
			args: args{
				ddl: string(input),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(mysqlParser)
			tableStruct, err := p.ParseMysqlDDL(tt.args.ddl)
			if err != nil {
				t.Errorf("ParseMysqlDDL() error = %v", err)
			}
			if err := tableStruct.GenerateFile(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		generateCode, err := os.ReadFile(filepath.Join(config.Cnf.OutputDir, "user.go"))
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}
		if !bytes.Equal(generateCode, golden) {
			t.Errorf("GenerateCode() is  = %v, want %v", generateCode, golden)
		}
	}
}
