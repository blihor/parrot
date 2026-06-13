package cmd

import (
	"fmt"

	"github.com/blihor/parrot/internal/generator"
	"github.com/spf13/cobra"
)

var (
	length         int
	includeUpper   bool
	includeDigits  bool
	includeSpecial bool
)

var cmdGeneratePassword = &cobra.Command{
	Use:   "gen [OPTIONS]",
	Short: "Generates random password",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		pass, err := generator.GeneratePassword(
			length,
			includeUpper,
			includeDigits,
			includeSpecial,
		)
		if err != nil {
			return err
		}

		fmt.Println(pass)

		return nil
	},
}

func init() {
	cmdGeneratePassword.Flags().IntVarP(&length, "length", "l", 16,
		"Set password length")
	cmdGeneratePassword.Flags().BoolVarP(&includeUpper, "upper", "u", false,
		"Include upper case letters in generator pool")
	cmdGeneratePassword.Flags().BoolVarP(&includeDigits, "digits", "d", false,
		"Include digits in generator pool")
	cmdGeneratePassword.Flags().BoolVarP(&includeSpecial, "special", "s", false,
		"Include special symbols in generator pool")
}
