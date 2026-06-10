package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
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
	rootCmd.AddCommand(cmdAddEntry)
	rootCmd.AddCommand(cmdDeleteEntry)
	rootCmd.AddCommand(cmdListEntry)
}
