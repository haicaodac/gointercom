package gointercom

import (
	"encoding/json"
	"log"
	"testing"
)

func TestConversationGetByID(t *testing.T) {
	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	jsonData, err := ConversationGetByID("20604517482")
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}

func TestConversationReply(t *testing.T) {
	config := GetConfig()
	config.SetToken("dG9rOjM0NWU4ZmE0X2Q5MGFfNDMxN184MWE4XzVjNDEzOTZjYjA5MjoxOjA=")

	data := make(map[string]interface{})

	data["admin_id"] = "2743023"
	data["body"] = `<p><a class="entity_mention" href="//app.intercom.io/apps/z11bwnw3/admin/2743023" rel="nofollow noopener noreferrer" target="_blank">Alana</a></p>`
	data["type"] = "admin"
	data["message_type"] = "note"

	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	ConversationReply("20604517482", string(jsonString))
}
