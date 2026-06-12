package cmd

import (
	"github.com/blihor/parrot/internal/auth"
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

		hs, key, err := auth.Authenticate(masterPassword, store)
		if err != nil {
			return err
		}

		vault, err := store.ReadVault(key)
		if err != nil {
			return err
		}

		entryName := args[0]
		vault.DeleteEntry(entryName)

		err = store.WriteStorage(key, hs, vault)
		if err != nil {
			return err
		}

		return nil
	},
}
