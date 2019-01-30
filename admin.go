package gointercom

import (
	"encoding/json"
	"log"
)

// AdminGetAdmins ..
func AdminGetAdmins() (AdminList, error) {
	url := "https://api.intercom.io/admins"
	byteData, err := Get(url)
	log.Println(string(byteData[:]))

	if err != nil {
		return AdminList{}, err
	}
	jsonData := AdminList{}

	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}
