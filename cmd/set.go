package cmd

import (
	"errors"

	"github.com/blihor/parrot/internal/auth"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/blihor/parrot/internal/storage"
	"github.com/spf13/cobra"
)

var cmdSetPassword = &cobra.Command{
	Use:   "set password",
	Short: "Sets new master password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()

		_, key, err := auth.Authenticate(masterPassword, store)
		if err != nil && !errors.Is(err, apperrors.ErrPasswordNotSet) {
			return err
		}

		newMasterPassword := args[0]

		return auth.SetPassword(newMasterPassword, key, store)
	},
}
