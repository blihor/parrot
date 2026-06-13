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
		var entryPass string

		if len(args) > 1 {
			entryPass = args[1]
		} else {
			entryPass, err = generator.GeneratePassword(
				v.GetInt("generator.length"),
				v.GetBool("generator.upper"),
				v.GetBool("generator.digits"),
				v.GetBool("generator.special"),
			)
		}

		if err != nil {
			return err
		}

		newEntry, err := entry.NewEntry(entryName, entryUrl, entryEmail, entryUsername, entryPass)
		if err != nil {
			return err
		}

		if err = vault.AddEntry(newEntry); err != nil {
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
	cmdAddEntry.Flags().StringVarP(&entryUsername, "username", "u", "", "Username for the entry")
	cmdAddEntry.Flags().StringVarP(&entryUrl, "url", "l", "", "Url associated with the entry")
	cmdAddEntry.Flags().StringVarP(&entryEmail, "email", "e", "", "Email associated with the entry")
}
