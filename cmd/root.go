package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
    userCfg DatabaseCfg

    rootCmd = &cobra.Command{
        Short: "Postgres mini cli to manage database",
    }
)

type DatabaseCfg struct {
	Usernname string
	Password  string
	Host      string
	Port      string
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&userCfg.Usernname, "username", "u", "", "Username for database. Default os usernaem")
	rootCmd.PersistentFlags().StringVarP(&userCfg.Password, "password", "p", "", "Password for database. If not specified you will be prompted for a password before connecting")
	rootCmd.PersistentFlags().StringVar(&userCfg.Host, "host", "localhost", "Database host")
	rootCmd.PersistentFlags().StringVar(&userCfg.Port, "port", "5432", "Database port")

    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(removeCmd)
    rootCmd.AddCommand(backupCmd)
    rootCmd.AddCommand(restoreCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func PrintErr(msg ...interface{}) {
    fmt.Fprint(os.Stderr, msg...)
    os.Exit(1)
}
