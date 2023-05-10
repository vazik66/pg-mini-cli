package cmd

import (
	"bytes"
	"fmt"
	"os"

	"os/exec"

	"github.com/spf13/cobra"
)

var (
	databaseName string
	restoreCmd   = &cobra.Command{
		Use:   "restore [backup file]",
		Short: "Restore database form backup",
		Long:  `Restore database from backup.
Recommended to create new database else all data will be erased!
Requires pg_restore installed, pg_restore version == database version`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pgrestore := exec.Command(
				"pg_restore",
                "--exit-on-error", 
                "--clean",
				"-U", userCfg.Usernname,
				"-h", userCfg.Host,
				"-p", userCfg.Port,
				"-d", databaseName,
				args[0],
			)

			pgrestore.Env = os.Environ()
			pgrestore.Env = append(pgrestore.Env, fmt.Sprintf("PGPASSWORD=%s", userCfg.Password))

			var stdErr bytes.Buffer
			pgrestore.Stderr = &stdErr

			fmt.Println("Restore started...")
			if err := pgrestore.Run(); err != nil {
				PrintErr(fmt.Sprintf("error restoring backup of \"%s\" database:\n%s", args[0], stdErr.String()))
			}
			fmt.Println("Database restored!")
		},
	}
)

func init() {
	restoreCmd.Flags().StringVarP(&databaseName, "dbname", "d", "", "Database name where the backup will be restored")
    _ = restoreCmd.MarkFlagRequired("dbname")
}
