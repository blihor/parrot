package cmd

import (
	"fmt"

	"github.com/blihor/parrot/internal/auth"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdListEntry = &cobra.Command{
	Use:   "list [NAME] ",
	Short: "List entry/entries",
	Long:  `List entry with the provided name or all entries if name was omitted`,
	RunE: WithConfig(func(cmd *cobra.Command, args []string, v *viper.Viper) error {
		store := storage.NewStorage(v.GetString("vault.filepath"))

		_, key, err := auth.Authenticate(masterPassword, store, v)
		if err != nil {
			return err
		}

		vault, err := store.ReadVault(key)
		if err != nil {
			return err
		}

		if len(args) > 0 {
			entryName := args[0]
			entry, err := vault.GetEntryByName(entryName)
			if err != nil {
				return err
			}

			fmt.Print(entry.String())
			return nil
		}

		entries := vault.GetAllEntries()
		for _, e := range entries {
			fmt.Print(e.String())
		}

		return nil
	}),
}
