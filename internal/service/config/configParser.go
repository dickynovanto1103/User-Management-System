package config

import (
	"encoding/json"
	"log"
	"os"
)

func openFile(filepath string, configFile *os.File) *json.Decoder {
	var err error
	configFile, err = os.Open(filepath)
	if err != nil {
		log.Println("error opening file:", err)
		return &json.Decoder{}
	}

	jsonParser := json.NewDecoder(configFile)
	return jsonParser
}

func LoadConfigRedis(filepath string) ConfigRedis {
	var configFile *os.File
	var config ConfigRedis

	jsonParser := openFile(filepath, configFile)
	defer configFile.Close()

	jsonParser.Decode(&config)

	return config
}

func LoadConfigDB(filepath string) ConfigDB {
	var configFile *os.File
	var config ConfigDB

	jsonParser := openFile(filepath, configFile)
	defer configFile.Close()

	jsonParser.Decode(&config)

	return config
}
