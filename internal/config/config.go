package config

import (
	"os"

	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName("parrot")
	v.AddConfigPath("$HOME/.config/parrot/")
	v.AddConfigPath("$HOME/.parrot")
	v.AddConfigPath("/etc/parrot/")

	setDefaults(v)

	err := v.ReadInConfig()
	if err != nil {
		return v, apperrors.ErrConfigNotFound
	}

	return v, nil
}

func setDefaults(v *viper.Viper) {
	// Generator
	v.SetDefault("generator.length", 16)
	v.SetDefault("generator.upper", true)
	v.SetDefault("generator.digits", true)
	v.SetDefault("generator.special", true)

	// Argon2d
	v.SetDefault("argon.time", 1)
	v.SetDefault("argon.mem", 64*1024)
	v.SetDefault("argon.threads", 4)
	v.SetDefault("argon.keylen", 32)
	v.SetDefault("argon.saltlen", 32)

	// Vault
	vaultFilePath := os.Expand("$HOME/.vault.json", func(s string) string {
		return os.Getenv(s)
	})
	v.SetDefault("vault.filepath", vaultFilePath)
}
