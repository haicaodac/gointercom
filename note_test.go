package gointercom

import (
	"log"
	"testing"
)

func TestNoteGetByUser(t *testing.T) {

	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	jsonData, err := NoteGetByUser()
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}
