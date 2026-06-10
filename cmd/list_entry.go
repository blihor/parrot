package cmd

import (
	"fmt"

	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
)

var cmdListEntry = &cobra.Command{
	Use:   "list [NAME] ",
	Short: "List entry/entries",
	Long:  `List entry with the provided name or all entries if name was omitted`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		vault, err := store.ReadVault([]byte(""))
		if err != nil {
			return err
		}

		if len(args) > 0 {
			entryName := args[1]
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
	},
}

func init() {
}
