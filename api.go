package main

type JournalApi struct {
	store JournalEntryStore
}

func (a JournalApi) CreateJournalEntry(content []byte) (JournalEntry, error) {
	entry := JournalEntry{Contents: content}
	j, err := a.store.Save(entry)
	return j, err
}

func (a JournalApi) GetJournalEntry(id uint) (JournalEntry, error) {
	entry, err := a.store.Get(id)
	return entry, err
}

func (a JournalApi) ListJournalEntries() []JournalEntry {
	entries := a.store.List()
	return entries
}
