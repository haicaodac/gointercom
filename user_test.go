package gointercom

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestUserGetUserByID(t *testing.T) {
	config := GetConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INTERCOM_TOKEN")
	config.SetToken(token)

	jsonData, err := UserGetUserByID("5c3f1af67e46af2a70214a51")
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}

func TestUserGetUsers(t *testing.T) {
	config := GetConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INTERCOM_TOKEN")
	config.SetToken(token)

	jsonData, err := UserGetUsers()
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}
