package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms
func loadBookworms(filePath string) ([]Bookworm, error){
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	
	// Decode the file and store the content in the variable bookworms
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}


// findCommonBooks returns books that are on more than one bookworm's shelf
func findCommonBooks(bookworms []Bookworm) []Book {
	return nil
}