package cmd

import (
	"fmt"
	"os"

	"github.com/gangming/sql2struct/config"
	"github.com/gangming/sql2struct/internal/driver"
	"github.com/gangming/sql2struct/internal/infra"
	"github.com/gangming/sql2struct/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sql2struct",
	Short: "sql2struct is a tool for generating golang struct from mysql/postgresql database",
	Long: `sql2struct is a tool for generating golang struct from mysql/postgresql database.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if config.Cnf.DSN == "" {
			utils.PrintRed("dsn is empty")
			_ = cmd.Help()
			os.Exit(1)
		}
		driverName := utils.GetDriverName(config.Cnf.DSN)
		infra.InitDB(driverName, config.Cnf.DSN)

		err := driver.NewSqlDriverGenerator(driverName).Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() {
		if err := recover(); err != nil {
			utils.PrintRedf("error occur %v", err)
		}
	}()
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

	rootCmd.PersistentFlags().StringVar(&config.Cnf.DSN, "dsn", "", "database dsn string (eg: root:123456@tcp(localhost:3306)/test?charset=utf8)")
	_ = rootCmd.MarkFlagRequired("dsn")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.DBTag, "dbtag", "g", "gorm", "db tag. default: gorm")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.TablePrefix, "prefix", "p", "", "table prefixed with the table name")
	rootCmd.PersistentFlags().BoolVarP(&config.Cnf.WithJsonTag, "with_json_tag", "j", true, "with json tag. default: true")
	rootCmd.PersistentFlags().BoolVarP(&config.Cnf.WithDefaultValue, "with_default_value", "", false, "with db default value. default: false")
	rootCmd.PersistentFlags().StringVarP(&config.Cnf.OutputDir, "output_dir", "o", "./model", "output dir. default: ./model")
	rootCmd.PersistentFlags().StringArrayVarP(&config.Cnf.Tables, "tables", "t", nil, "Need to generate tables, default is all tables. (eg: -t table1 -t table2)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
