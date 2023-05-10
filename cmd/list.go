package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vazik66/pg-mini-cli/internal"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List databases",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := internal.NewClient(
				userCfg.Usernname,
				userCfg.Password,
				userCfg.Host,
				userCfg.Port,
			)

			if err != nil {
                PrintErr(fmt.Sprintf("Could not connect to database: %v", err))
			}
			defer client.Conn.Close(context.Background())

			databases, err := client.ListDatabases()
			if err != nil {
                PrintErr(fmt.Sprintf("Could not get list of databases: %v", err))
			}
            formatted_out := strings.Join(databases, "\n")
            fmt.Println(formatted_out)
		},
	}
)
