package cmd

import (
	"github.com/blihor/parrot/internal/entry"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
)

var (
	entryUsername string
	entryUrl      string
	entryEmail    string
)

var cmdAddEntry = &cobra.Command{
	Use:   "add name password [OPTIONS]",
	Short: "Add new entry to the vault",
	Long: `Add a new entry of a provided name. All fields besides name and 
        password are optional
        `,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		vault, err := store.ReadVault([]byte(""))
		if err != nil {
			return err
		}

		entryName := args[0]
		entryPass := args[1]
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
