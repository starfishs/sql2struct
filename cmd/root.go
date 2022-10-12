package cmd

import (
	"fmt"
	"github.com/gangming/sql2struct/config"
	"github.com/gangming/sql2struct/internal/infra"
	mysqlparser "github.com/gangming/sql2struct/internal/mysql"
	"github.com/gangming/sql2struct/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sql2struct",
	Short: "sql2struct is a tool for generating golang struct from mysql database",
	Long: `sql2struct is a tool for generating golang struct from mysql database.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if config.Cnf.DSN == "" {
			utils.PrintRed("dsn is empty")
			cmd.Help()
			os.Exit(1)
		}
		infra.InitDBMysql(config.Cnf.DSN)
		err := mysqlparser.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sql2struct.yaml)")

	rootCmd.PersistentFlags().StringVar(&config.Cnf.DSN, "dsn", "", "database dsn string (eg: root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local)")
	rootCmd.MarkFlagRequired("dsn")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.DBTag, "dbtag", "g", "", "db tag. default: gorm")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.TablePrefix, "prefix", "p", "", "table prefixed with the table name")
	rootCmd.PersistentFlags().BoolVarP(&config.Cnf.WithJsonTag, "with_json_tag", "j", true, "with json tag. default: true")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.OutputDir, "output_dir", "o", "./model", "output dir. default: ./model")
	rootCmd.PersistentFlags().StringArrayVarP(&config.Cnf.Tables, "tables", "t", []string{""}, "Need to generate tables, default is all tables. (eg: -t table1 -t table2)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
