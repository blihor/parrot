package auth

import (
	"errors"
	"fmt"
	"io/fs"
	"syscall"

	"github.com/blihor/parrot/internal/crypto"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/blihor/parrot/internal/storage"
	"golang.org/x/term"
)

const maxAttempts = 5

type (
	Argon2idHash = crypto.Argon2idHash
	HashSalt     = crypto.HashSalt
)

// ValidatePassword validates masterPassword and, if successfull, returns
// encryption key
func ValidatePassword(masterPassword string, hs *HashSalt) ([]byte, error) {
	argon2 := crypto.NewArgon2idHash(1, 64*1024, 4, 32, 32)

	encryptionKey, err := argon2.Compare(hs.Hash, hs.Salt, []byte(masterPassword))
	if err != nil {
		return nil, err
	}

	return encryptionKey, nil
}

func HashPassword(masterPassword string) (*HashSalt, error) {
	argon2 := crypto.NewArgon2idHash(1, 64*1024, 4, 32, 32)

	return argon2.GenerateHash([]byte(masterPassword), nil)
}

func SetPassword(newMasterPassword string, key []byte, store *storage.Storage) error {
	vault, err := store.ReadVault(key)
	if err != nil {
		return err
	}

	newHs, err := HashPassword(newMasterPassword)
	if err != nil {
		return err
	}

	// Get a new key from hash
	hashLen := len(newHs.Hash) / 2
	newKey := newHs.Hash[hashLen:]

	// Strip hash of encryption key
	newHs.Hash = newHs.Hash[:hashLen]

	return store.WriteVaultAndHashSalt(newKey, newHs, vault)
}

func PromptForPassword() (string, error) {
	fmt.Print("Enter master password: ")

	masterPassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}
	if len(masterPassword) == 0 {
		return "", apperrors.ErrEmptyPassword
	}

	fmt.Println()

	return string(masterPassword), nil
}

// Authenticate reads hash and salt from file, compares it with actually generated
// hash out of masterPassword and returns hashSalt struct with the hash and the
// salt, and encryption key. If file doesn't exists it returns ErrPasswordNotSet.
// User will be prompted for password if one wasn't provided
func Authenticate(masterPassword string, store *storage.Storage) (*crypto.HashSalt, []byte, error) {
	var hs *storage.HashSalt
	var err error

	hs, err = store.ReadHashSalt()
	if err != nil {
		if errors.Is(err, apperrors.ErrUnmarshalFailed) || errors.Is(err, fs.ErrNotExist) {
			return nil, nil, apperrors.ErrPasswordNotSet
		} else {
			return nil, nil, err
		}
	}

	failedAttempts := 0
	var encryptionKey []byte

	for failedAttempts < maxAttempts {
		if len(masterPassword) == 0 {
			masterPassword, err = PromptForPassword()
			if err != nil {
				if errors.Is(err, apperrors.ErrEmptyPassword) {
					fmt.Printf("Password is empty. Please provide actual password\n")
					continue
				}

				return nil, nil, err
			}
		}

		encryptionKey, err = ValidatePassword(masterPassword, hs)
		if err != nil {
			if errors.Is(err, apperrors.ErrHashNotEqual) {
				masterPassword = ""
				failedAttempts++

				fmt.Printf("%s\n\n", err.Error())
				fmt.Printf("Attempts left: %d\n\n", maxAttempts-failedAttempts)

				continue
			}

			return nil, nil, err
		}

		break
	}

	if failedAttempts == maxAttempts {
		return nil, nil, apperrors.ErrAttemptsFailed
	}

	return hs, encryptionKey, nil
}
