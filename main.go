package main

func main() {
	// store := GetJournalEntryStore("memory")
	// api := &JournalApi{store: store}
	// fmt.Println(api.ListJournalEntries())
	// _, err := api.CreateJournalEntry([]byte("hello"))
	// if err != nil {
	// 	errors.Errorf("Failed, %s", err)
	// }
	// _, err = api.CreateJournalEntry([]byte("boo yah"))
	// if err != nil {
	// 	errors.Errorf("Failed, %s", err)
	// }
	// fmt.Println(api.ListJournalEntries())
	server := NewServer()
	server.StartServer()
}
