package vault

import (
	"github.com/blihor/parrot/internal/entry"
	apperrors "github.com/blihor/parrot/internal/errors"
)

type Entry = entry.Entry

type Vault struct {
	EntryMap map[string]*Entry
}

func NewVault() *Vault {
	return &Vault{
		EntryMap: map[string]*Entry{},
	}
}

func (v *Vault) AddEntry(entry *Entry) error {
	_, exists := v.EntryMap[entry.Name]
	if exists {
		return apperrors.ErrEntryNameTaken
	}

	v.EntryMap[entry.Name] = entry
	return nil
}

func (v *Vault) EditEntry(name string, updateData *Entry) error {
	entry, err := v.GetEntryByName(name)
	if err != nil {
		return err
	}

	if updateData.Name != "_" {
		entry.Name = updateData.Name
	}
	if updateData.Password != "_" {
		entry.Password = updateData.Password
	}
	if updateData.Url != "_" {
		entry.Url = updateData.Url
	}
	if updateData.Email != "_" {
		entry.Email = updateData.Email
	}
	if updateData.Username != "_" {
		entry.Username = updateData.Username
	}

	return nil
}

func (v *Vault) DeleteEntry(name string) {
	delete(v.EntryMap, name)
}

func (v *Vault) GetEntryByName(name string) (*Entry, error) {
	entry, exists := v.EntryMap[name]
	if !exists {
		return nil, apperrors.ErrEntryNotFound
	}

	return entry, nil
}

func (v *Vault) GetAllEntries() []*Entry {
	entries := make([]*Entry, 0, len(v.EntryMap))

	for _, e := range v.EntryMap {
		entries = append(entries, e)
	}

	return entries
}
