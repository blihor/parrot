package apperrors

import "errors"

var (
	ErrNoEntryNameProvided     = errors.New("entry name is missing")
	ErrNoEntryPasswordProvided = errors.New("entry password is missing")
	// TODO make error msg more clear for user to know that password is incorrect
	ErrHashNotEqual    = errors.New("hash does not match")
	ErrShortCipher     = errors.New("cipher is to short")
	ErrEntryNotFound   = errors.New("entry not found")
	ErrEntryNameTaken  = errors.New("provided name is already taken")
	ErrUnmarshalFailed = errors.New("bad json format. Unmarhsal operation failed")
	ErrEmptyPassword   = errors.New("master password could not be empty")
	ErrPasswordNotSet  = errors.New("password is not set. use set command with default password to set one")
	ErrAttemptsFailed  = errors.New("too many attempts")
)
