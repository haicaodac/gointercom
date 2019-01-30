package gointercom

import (
	"encoding/json"
)

// ConversationGetByID ...
func ConversationGetByID(id string) (IntercomConversation, error) {
	url := "https://api.intercom.io/conversations/" + id
	byteData, err := Get(url)
	if err != nil {
		return IntercomConversation{}, err
	}

	jsonData := IntercomConversation{}

	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}
