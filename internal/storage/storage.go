package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/blihor/parrot/internal/crypto"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/blihor/parrot/internal/vault"
)

type (
	Vault    = vault.Vault
	HashSalt = crypto.HashSalt
)

type (
	Storage struct {
		VaultFilePath string
	}

	storageData struct {
		HashSalt       *HashSalt
		EncryptedVault []byte
	}
)

func NewStorage(filepath string) *Storage {
	return &Storage{
		VaultFilePath: filepath,
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

	storageData := &storageData{}
	if err := json.Unmarshal(jsonData, storageData); err != nil {
		return nil, err
	}

	vaultJSON, err := crypto.Decrypt(key, storageData.EncryptedVault)
	if err != nil {
		return nil, err
	}

	vault := &Vault{}
	if err := json.Unmarshal(vaultJSON, vault); err != nil {
		return nil, err
	}

	return vault, nil
}

func (s *Storage) WriteStorage(key []byte, hashSalt *HashSalt, v *Vault) error {
	vaultJSON, err := json.Marshal(v)
	if err != nil {
		return err
	}

	encryptedVault, err := crypto.Encrypt(key, vaultJSON)
	if err != nil {
		return err
	}

	storageData := &storageData{
		HashSalt:       hashSalt,
		EncryptedVault: encryptedVault,
	}

	dataJSON, err := json.Marshal(storageData)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.VaultFilePath, dataJSON, 0o751); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ReadHashSalt() (*HashSalt, error) {
	jsonData, err := os.ReadFile(s.VaultFilePath)
	if err != nil {
		return nil, err
	}

	storageData := &storageData{}
	if err := json.Unmarshal(jsonData, storageData); err != nil {
		return nil, apperrors.ErrUnmarshalFailed
	}

	return storageData.HashSalt, nil
}

func (s *Storage) DeleteStorage() {
	fmt.Println("Deleting storage...")
	os.Remove(s.VaultFilePath)
}
