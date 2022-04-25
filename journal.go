package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

var (
	ErrorSavingJournalEntry  = errors.New("Error saving journal entry")
	ErrorGettingJournalEntry = errors.New("Error getting journal entry")
)

type JournalEntry struct {
	Id       uint
	Contents string
}

func (e JournalEntry) String() string {
	return fmt.Sprintf("%d: %s", e.Id, e.Contents)
}

type JournalEntryStore interface {
	Save(JournalEntry) (JournalEntry, error)
	Get(uint) (JournalEntry, error)
	List() []JournalEntry
}

type JournalEntryMemoryStore struct {
	entries []JournalEntry
}

func (s *JournalEntryMemoryStore) Save(j JournalEntry) (JournalEntry, error) {
	var maxId uint
	if len(s.entries) == 0 {
		maxId = 0
	} else {
		maxId = s.entries[len(s.entries)-1].Id
	}
	savedEntry := JournalEntry{
		Id:       maxId + 1,
		Contents: j.Contents,
	}
	s.entries = append(s.entries, savedEntry)
	return savedEntry, nil
}

func (s *JournalEntryMemoryStore) Get(id uint) (JournalEntry, error) {
	var entry JournalEntry
	var found bool = false
	for _, j := range s.entries {
		if j.Id == id {
			entry = j
			found = true
			break
		}
	}
	if found {
		return entry, nil
	} else {
		return entry, ErrorGettingJournalEntry
	}
}

func (s *JournalEntryMemoryStore) List() []JournalEntry {
	return s.entries
}

type JournalEntryFileStore struct {
	filepath string
}

func (s *JournalEntryFileStore) Save(j JournalEntry) (JournalEntry, error) {
	filename := fmt.Sprintf("%d.txt", j.Id)
	err := os.WriteFile(filename, []byte(j.Contents), 640)
	if err != nil {
		err = errors.Wrap(ErrorSavingJournalEntry, err.Error())
	}
	return j, err
}

func GetJournalEntryStore(mode string) JournalEntryStore {
	var j JournalEntryStore
	switch mode {
	case "memory":
		j = &JournalEntryMemoryStore{entries: make([]JournalEntry, 0)}
		// case "file":
		// 	j = &JournalEntryFileStore{filepath: "./data"}
	}
	return j
}
