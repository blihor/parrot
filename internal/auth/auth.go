package auth

import (
	"errors"
	"io/fs"

	"github.com/blihor/parrot/internal/crypto"
	apperrors "github.com/blihor/parrot/internal/errors"
	"github.com/blihor/parrot/internal/storage"
)

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

// Authenticate reads hash and salt from file, compares it with actually generated
// hash out of masterPassword and returns hashSalt struct with the hash and the
// salt, and encryption key. If file doesn't exists it hashes masterPassword
// and returns same results(hashSalf and key) as in former case, but without any
// validation since it is the first time password is set
func Authenticate(masterPassword string, store *storage.Storage) (*crypto.HashSalt, []byte, error) {
	var hs *storage.HashSalt
	var err error

	hs, err = store.ReadHashSalt()
	if err != nil {
		if errors.Is(err, apperrors.ErrUnmarshalFailed) || errors.Is(err, fs.ErrNotExist) {
			hs, err = HashPassword(masterPassword)
			if err != nil {
				return nil, nil, err
			}

			// Strip hash of encryption key
			hashLen := len(hs.Hash) / 2
			encryptionKey := hs.Hash[hashLen:]
			hs.Hash = hs.Hash[:hashLen]

			return hs, encryptionKey, nil
		} else {
			return nil, nil, err
		}
	}

	encryptionKey, err := ValidatePassword(masterPassword, hs)
	if err != nil {
		return nil, nil, err
	}

	return hs, encryptionKey, nil
}
