package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// AuthInfo represents the authentication information from the config file
type AuthInfo struct {
	Password string `json:"AUTH_PASSWORD"`
	Username string `json:"AUTH_USERNAME"`
}

// DbInfo represents the database connection information from the config file
type DbInfo struct {
	Host     string `json:"DB_HOST"`
	Name     string `json:"DB_NAME"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	User     string `json:"DB_USER"`
}

// Read reads and parses the configuration file and returns AuthInfo and DbInfo structs
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

	// Decode AuthInfo from the configuration file
	var auth AuthInfo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&auth)
	if err != nil {
		fmt.Println("Error parsing AuthInfo:", err)
		return AuthInfo{}, DbInfo{}, err
	}

	// Reset file pointer to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return AuthInfo{}, DbInfo{}, err
	}

	// Decode DbInfo from the configuration file
	var db DbInfo
	err = decoder.Decode(&db)
	if err != nil {
		fmt.Println("Error parsing DBInfo:", err)
		return AuthInfo{}, DbInfo{}, err
	}
	return auth, db, nil
}
