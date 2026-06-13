package cmd

import (
	"github.com/blihor/parrot/internal/auth"
	"github.com/blihor/parrot/internal/entry"
	"github.com/blihor/parrot/internal/generator"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	entryPassword    string
	entryNameUpdated string
	useGenerator     bool
)

var cmdEditEntry = &cobra.Command{
	Use:   "edit name [OPTIONS]",
	Short: "Edit existsing entry",
	Long: `Edit entry by name. All options are optional, but at least one
        should be provided.

        Password it set manually via --password or -p flag OR generates automatically
        via --gen or -g flag, but not both.`,
	Args:             cobra.ExactArgs(1),
	TraverseChildren: true,
	RunE: WithConfig(func(cmd *cobra.Command, args []string, v *viper.Viper) error {
		store := storage.NewStorage(v.GetString("vault.filepath"))

		hs, key, err := auth.Authenticate(masterPassword, store, v)
		if err != nil {
			return err
		}

		vault, err := store.ReadVault(key)
		if err != nil {
			return err
		}

		entryName := args[0]

		if useGenerator {
			entryPassword, err = generator.GeneratePassword(
				v.GetInt("generator.length"),
				v.GetBool("generator.upper"),
				v.GetBool("generator.digits"),
				v.GetBool("generator.special"),
			)
		}
		if err != nil {
			return err
		}

		updateData, err := entry.NewEntry(entryNameUpdated, entryUrl, entryEmail, entryUsername, entryPassword)
		if err != nil {
			return err
		}

		if err = vault.EditEntry(entryName, updateData); err != nil {
			return err
		}

		err = store.WriteStorage(key, hs, vault)
		if err != nil {
			return err
		}

		return nil
	}),
}

func init() {
	cmdEditEntry.Flags().StringVarP(&entryNameUpdated, "name", "n", "_", "Name of the entry")
	cmdEditEntry.Flags().StringVarP(&entryUsername, "username", "u", "_", "Username for the entry")
	cmdEditEntry.Flags().StringVarP(&entryUrl, "url", "l", "_", "Url associated with the entry")
	cmdEditEntry.Flags().StringVarP(&entryEmail, "email", "e", "_", "Email associated with the entry")
	cmdEditEntry.Flags().StringVarP(&entryPassword, "password", "p", "_", "Password associated with the entry")
	cmdEditEntry.Flags().BoolVarP(&useGenerator, "gen", "g", false, "Use generator for new password")
	cmdEditEntry.MarkFlagsOneRequired("name", "username", "url", "email", "password")
	cmdEditEntry.MarkFlagsMutuallyExclusive("password", "gen")
}
