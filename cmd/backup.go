package cmd

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"os/exec"

	"github.com/spf13/cobra"
)

var (
	backupFileName string
	backupCmd      = &cobra.Command{
		Use:   "backup",
		Short: "Create backup of database",
		Long:  "Create backup of database. Requires pg_dump installed, pg_dump version == database version",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if backupFileName == "" {
				backupFileName = fmt.Sprintf(
					"%s_%s.dump",
					args[0],
					time.Now().UTC().Format(time.DateOnly),
				)
			}

			pgdump := exec.Command(
				"pg_dump",
				"-U", userCfg.Usernname,
				"-h", userCfg.Host,
				"-p", userCfg.Port,
				"-Fc",
				args[0],
				"--file", backupFileName,
			)
			pgdump.Env = os.Environ()
			pgdump.Env = append(pgdump.Env, fmt.Sprintf("PGPASSWORD=%s", userCfg.Password))

			var stdErr bytes.Buffer
			pgdump.Stderr = &stdErr

			fmt.Println("Backup started...")
			if err := pgdump.Run(); err != nil {
				PrintErr(fmt.Sprintf("error creating backup of \"%s\" database:\n%s", args[0], stdErr.String()))
			}
			fmt.Printf("Backup created! -> %s\n", backupFileName)
		},
	}
)

func init() {
	backupCmd.Flags().StringVarP(&backupFileName, "file", "f", "", "Backup filename")
}
