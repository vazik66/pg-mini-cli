package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vazik66/pg-mini-cli/internal"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove [dbname1 dbname2...]",
		Short: "Remove databases",
		Args:  cobra.MinimumNArgs(1),
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
			availableDatabases, err := client.ListDatabases()
			if err != nil {
				PrintErr(fmt.Sprintf("Could not get databases: %v", err))
			}

			// Check databases exist
			databasesMap := make(map[string]bool, len(availableDatabases))
			for _, database := range availableDatabases {
				databasesMap[database] = true
			}

			for _, database := range args {
				if _, ok := databasesMap[database]; !ok {
					PrintErr(fmt.Sprintf("Database \"%s\" does not exist\n", database))
					os.Exit(1)
				}
			}

			// Remove databases
			for _, database := range args {
				err := client.DeleteDatabase(database)
				if err != nil {
					fmt.Printf("Error removing \"%s\" database\n", database)
				} else {
					fmt.Printf("Database \"%s\" removed!\n", database)
				}
			}

		},
	}
)
