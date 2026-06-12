package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var masterPassword string

var rootCmd = &cobra.Command{
	Use:   "parrot",
	Short: "CLI password manager",
	Long: `Parrot is a CLI for managing your passwords that comes with 
        built-in configurable password generator. It operates on entries, which
        are combination of the name or url, email or username and password.
        `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&masterPassword, "master", "m", "", "Master password to access vault")
	rootCmd.MarkPersistentFlagRequired("master")
	rootCmd.AddCommand(cmdAddEntry)
	rootCmd.AddCommand(cmdDeleteEntry)
	rootCmd.AddCommand(cmdListEntry)
	rootCmd.AddCommand(cmdGeneratePassword)
	rootCmd.AddCommand(cmdSetPassword)
}
