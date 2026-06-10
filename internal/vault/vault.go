package vault

import (
	"github.com/blihor/parrot/internal/entry"
	"github.com/blihor/parrot/internal/errors"
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
		return errors.ErrEntryNameTaken
	}

	v.EntryMap[entry.Name] = entry
	return nil
}

func (v *Vault) DeleteEntry(name string) {
	delete(v.EntryMap, name)
}

func (v *Vault) GetEntryByName(name string) (*Entry, error) {
	entry, exists := v.EntryMap[name]
	if !exists {
		return nil, errors.ErrEntryNotFound
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
