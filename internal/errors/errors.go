package apperrors

import "errors"

var (
	ErrNoEntryNameProvided     = errors.New("entry name is missing")
	ErrNoEntryPasswordProvided = errors.New("entry password is missing")
	ErrHashNotEqual            = errors.New("hash does not match")
	ErrShortCipher             = errors.New("cipher is to short")
	ErrEntryNotFound           = errors.New("entry not found")
	ErrEntryNameTaken          = errors.New("provided name is already taken")
	ErrUnmarshalFailed         = errors.New("bad json format. Unmarhsal operation failed")
)
