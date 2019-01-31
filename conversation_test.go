package gointercom

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConversationGetByID(t *testing.T) {
	config := GetConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INTERCOM_TOKEN")
	config.SetToken(token)

	jsonData, err := ConversationGetByID("20604517482")
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonData)
}

func TestConversationReply(t *testing.T) {
	config := GetConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INTERCOM_TOKEN")
	config.SetToken(token)

	data := make(map[string]interface{})

	data["admin_id"] = "2743023"
	data["body"] = `<p><a class="entity_mention" href="//app.intercom.io/apps/z11bwnw3/admin/2928092" rel="nofollow noopener noreferrer" target="_blank">Minh Nora</a> 😛😛😛 Content</p>`
	data["type"] = "admin"
	data["message_type"] = "note"

	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	ConversationReply("20633013706", jsonByte)
}
