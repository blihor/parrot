package cmd

import (
	"errors"
	"fmt"

	"github.com/blihor/parrot/internal/config"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type cobraFunc func(*cobra.Command, []string) error

func WithConfig(fn func(cmd *cobra.Command, args []string, v *viper.Viper) error) cobraFunc {
	return func(cmd *cobra.Command, args []string) error {
		v, err := config.LoadConfig()
		if err != nil {
			if errors.Is(err, apperrors.ErrConfigNotFound) {
				fmt.Println("\n***Config file not found. Fallback to defaults***\n\n")
			} else {
				return err
			}
		} else {
			fmt.Printf("\n***Config file used: %s***\n\n", v.ConfigFileUsed())
		}

		err = fn(cmd, args, v)
		if err != nil {
			return err
		}

		return nil
	}
}
