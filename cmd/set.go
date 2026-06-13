package cmd

import (
	"errors"

	"github.com/blihor/parrot/internal/auth"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdSetPassword = &cobra.Command{
	Use:   "set password",
	Short: "Sets new master password",
	Args:  cobra.ExactArgs(1),
	RunE: WithConfig(func(cmd *cobra.Command, args []string, v *viper.Viper) error {
		store := storage.NewStorage(v.GetString("vault.filepath"))

		_, key, err := auth.Authenticate(masterPassword, store, v)
		if err != nil && !errors.Is(err, apperrors.ErrPasswordNotSet) {
			return err
		}

		newMasterPassword := args[0]

		return auth.SetPassword(newMasterPassword, key, store, v)
	}),
}
