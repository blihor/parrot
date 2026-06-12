package entry

import (
	"fmt"
	"strings"

	apperrors "github.com/blihor/parrot/internal/errors"
)

type Entry struct {
	Name     string `json:"name"`
	Url      string `json:"url,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

func (e *Entry) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("Name: %s\n", e.Name))

	if e.Url != "" {
		lines = append(lines, fmt.Sprintf("Url: %s\n", e.Url))
	}
	if e.Email != "" {
		lines = append(lines, fmt.Sprintf("Email: %s\n", e.Email))
	}
	if e.Username != "" {
		lines = append(lines, fmt.Sprintf("Username: %s\n", e.Username))
	}

	lines = append(lines, fmt.Sprintf("Password: %s\n", e.Password))

	return strings.Join(lines, "")
}

func NewEntry(
	name string,
	url string,
	email string,
	username string,
	password string,
) (*Entry, error) {
	if name == "" {
		return nil, apperrors.ErrNoEntryNameProvided
	}

	if password == "" {
		return nil, apperrors.ErrNoEntryPasswordProvided
	}

	return &Entry{
		Name:     name,
		Url:      url,
		Email:    email,
		Username: username,
		Password: password,
	}, nil
}
