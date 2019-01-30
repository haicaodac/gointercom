package gointercom

import (
	"encoding/json"
)

// NoteGetByUser ..
func NoteGetByUser() (IntercomNotes, error) {
	url := "https://api.intercom.io/notes?email=getshoplaunch@gmail.com"
	byteData, err := Get(url)
	if err != nil {
		return IntercomNotes{}, err
	}
	jsonData := IntercomNotes{}

	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}
