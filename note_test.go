package gointercom

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNoteGetByUser(t *testing.T) {

	config := GetConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INTERCOM_TOKEN")
	config.SetToken(token)

	jsonData, err := NoteGetByUser()
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}
