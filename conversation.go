package gointercom

import (
	"encoding/json"
	"log"
)

/*
ConversationGetByID ...
	@param: id string // conversation_id
*/
func ConversationGetByID(id string) (IntercomConversation, error) {
	url := "https://api.intercom.io/conversations/" + id
	byteData, err := Get(url)
	log.Println(string(byteData[:]))
	if err != nil {
		return IntercomConversation{}, err
	}

	jsonData := IntercomConversation{}

	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}

/*
ConversationReply ...
	@param: id string // conversation_id
			data string // json data
*/
func ConversationReply(id string, data string) error {
	url := "https://api.intercom.io/conversations/" + id + "/reply"
	byteData, err := Post(url, data)
	if err != nil {
		return err
	}

	log.Println(string(byteData))
	return nil
}
