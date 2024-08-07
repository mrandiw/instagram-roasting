package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ApiKeyGemini string `json:"API_KEY_GEMINI"`
	TextRoast    string `json:"TEXT_ROAST"`
}

func GetConfig() Config {
	// Read the json file
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config.json file: %v", err)
	}

	// Unmarshal the json file into a Config struct
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling json: %v", err)
	}

	return config
}
