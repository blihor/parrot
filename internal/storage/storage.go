package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/blihor/parrot/internal/vault"
)

var vaultFilePath = os.Expand("$HOME/.vault.json", func(s string) string {
	if s == "HOME" {
		home, _ := os.UserHomeDir()
		return home
	}

	return os.Getenv(s)
})

type Vault = vault.Vault

type Storage struct {
	VaultFilePath string
}

func NewStorage() *Storage {
	return &Storage{
		VaultFilePath: vaultFilePath,
	}
}

func (s *Storage) ReadVault(key []byte) (*Vault, error) {
	jsonData, err := os.ReadFile(s.VaultFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return vault.NewVault(), nil
		} else {
			return nil, err
		}
	}

	// jsonData, err := crypto.Decrypt(key, data)
	// if err != nil {
	// 	return nil, err
	// }

	vault := &Vault{}
	if err := json.Unmarshal(jsonData, vault); err != nil {
		return nil, err
	}

	return vault, nil
}

func (s *Storage) WriteVault(key []byte, v *Vault) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// encryptedData, err := crypto.Encrypt(key, jsonData)
	// if err != nil {
	// 	return err
	// }

	if err := os.WriteFile(s.VaultFilePath, jsonData, 0o751); err != nil {
		return err
	}

	return nil
}
