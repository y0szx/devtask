package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AuthInfo struct {
	Password string `json:"AUTH_PASSWORD"`
	Username string `json:"AUTH_USERNAME"`
}

type DbInfo struct {
	Host     string `json:"DB_HOST"`
	Name     string `json:"DB_NAME"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	User     string `json:"DB_USER"`
}

func Read() (AuthInfo, DbInfo, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return AuthInfo{}, DbInfo{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var auth AuthInfo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&auth)
	if err != nil {
		fmt.Println("Error parsing AuthInfo:", err)
		return AuthInfo{}, DbInfo{}, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return AuthInfo{}, DbInfo{}, err
	}

	var db DbInfo
	err = decoder.Decode(&db)
	if err != nil {
		fmt.Println("Error parsing DBInfo:", err)
		return AuthInfo{}, DbInfo{}, err
	}
	return auth, db, nil
}
