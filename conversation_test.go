package gointercom

import (
	"log"
	"testing"
)

func TestConversationGetByID(t *testing.T) {
	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	jsonData, err := ConversationGetByID("20497849997")
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}
