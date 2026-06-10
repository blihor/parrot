package cmd

import (
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
)

var cmdDeleteEntry = &cobra.Command{
	Use:   "delete name ",
	Short: "delete entry",
	Long:  `delete entry with the provided name`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		vault, err := store.ReadVault([]byte(""))
		if err != nil {
			return err
		}

		entryName := args[0]
		vault.DeleteEntry(entryName)

		store.WriteVault([]byte(""), vault)
		return nil
	},
}

func init() {
}
