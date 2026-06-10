package generator

import (
	"crypto/rand"
	"math/big"
)

// TODO: configure for max possible sequence of the same character

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func GeneratePassword(
	length int,
	includeUpper bool,
	includeDigits bool,
	includeSpecial bool,
) (string, error) {
	charPool := lowerChars

	if includeUpper {
		charPool += upperChars
	}
	if includeDigits {
		charPool += digitChars
	}
	if includeSpecial {
		charPool += specialChars
	}

	poolLength := big.NewInt(int64(len(charPool)))
	pass := make([]byte, length)

	for i := range length {
		randIdx, err := rand.Int(rand.Reader, poolLength)
		if err != nil {
			return "", err
		}

		pass[i] = charPool[randIdx.Int64()]
	}

	return string(pass), nil
}
