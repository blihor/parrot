package crypto

import (
	"bytes"
	"crypto/rand"

	apperrors "github.com/blihor/parrot/internal/errors"
	"golang.org/x/crypto/argon2"
)

type Argon2idHash struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

type HashSalt struct {
	Hash []byte
	Salt []byte
}

func NewArgon2idHash(
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
	saltLen uint32,
) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
		saltLen: saltLen,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// GenerateHash generates hash of size 2*keyLen, where first half is hash for
// auth and second is encryption/decryption key
func (a *Argon2idHash) GenerateHash(password []byte, salt []byte) (*HashSalt, error) {
	var err error

	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen*2)

	return &HashSalt{Hash: hash, Salt: salt}, nil
}

// Compare compares hash in arguments with the first half or actually generated
// hash from salf and password arguments. If compared successfully returns the
// second half for encryption/decryption purposes
func (a *Argon2idHash) Compare(hash []byte, salt []byte, password []byte) ([]byte, error) {
	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return nil, err
	}

	authHash := hashSalt.Hash[:a.keyLen]
	if !bytes.Equal(hash, authHash) {
		return nil, apperrors.ErrHashNotEqual
	}

	return hashSalt.Hash[a.keyLen:], nil
}
