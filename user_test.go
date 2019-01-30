package gointercom

import (
	"log"
	"testing"
)

func TestUserGetUserByID(t *testing.T) {
	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	jsonData, err := UserGetUserByID("5c3f1af67e46af2a70214a51")
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}

func TestUserGetUsers(t *testing.T) {
	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	jsonData, err := UserGetUsers()
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}
