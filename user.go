package gointercom

import (
	"encoding/json"
	"log"
)

// UserGetUserByID ..
func UserGetUserByID(id string) (IntercomUser, error) {
	url := "https://api.intercom.io/admins"
	byteData, err := Get(url)
	log.Println(string(byteData[:]))

	if err != nil {
		return IntercomUser{}, err
	}
	jsonData := IntercomUser{}

	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}

// UserGetUsers ..
func UserGetUsers() (IntercomUsers, error) {
	url := "https://api.intercom.io/users?email=" + "getshoplaunch@gmail.com"
	byteData, err := Get(url)
	// log.Println(string(byteData[:]))
	if err != nil {
		return IntercomUsers{}, err
	}
	jsonData := IntercomUsers{}
	json.Unmarshal(byteData, &jsonData)
	return jsonData, nil
}
