package errors

import "errors"

var (
	ErrNoEntryNameProvided     = errors.New("entry name is missing")
	ErrNoEntryPasswordProvided = errors.New("entry password is missing")
	ErrHashNotEqual            = errors.New("hash does not match")
	ErrShortCipher             = errors.New("cipher is to short")
	ErrEntryNotFound           = errors.New("entry not found")
)
