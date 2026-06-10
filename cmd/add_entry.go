package cmd

import (
	"github.com/blihor/parrot/internal/entry"
	"github.com/blihor/parrot/internal/generator"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
)

var (
	entryUsername string
	entryUrl      string
	entryEmail    string
)

var cmdAddEntry = &cobra.Command{
	Use:   "add name [password] [OPTIONS]",
	Short: "Add new entry to the vault",
	Long: `Add a new entry of a provided name. All fields besides name are 
        optional. If password is not provided it will be generated randomly.
        `,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		vault, err := store.ReadVault([]byte(""))
		if err != nil {
			return err
		}

		entryName := args[0]
		var entryPass string

		if len(args) > 1 {
			entryPass = args[1]
		} else {
			// TODO: configure generator
			entryPass, err = generator.GeneratePassword(16, true, true, true)
		}

		if err != nil {
			return err
		}

		newEntry, err := entry.NewEntry(entryName, entryUrl, entryEmail, entryUsername, entryPass)
		if err != nil {
			return err
		}

		vault.AddEntry(newEntry)

		store.WriteVault([]byte(""), vault)
		return nil
	},
}

func init() {
	cmdAddEntry.Flags().StringVarP(&entryUsername, "username", "u", "", "Username for the entry")
	cmdAddEntry.Flags().StringVarP(&entryUrl, "url", "l", "", "Url associated with the entry")
	cmdAddEntry.Flags().StringVarP(&entryEmail, "email", "e", "", "Email associated with the entry")
}
