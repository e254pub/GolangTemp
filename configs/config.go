package configs

import (
	"github.com/joho/godotenv"
	"log"
	"main/pkg/convert"
	"main/pkg/utils"
	"path/filepath"
)

func LoadEnv() {
	currentDirectory, ok := utils.GetFileDirectoryToCaller()
	if !ok {
		panic("Failed to load configuration: Failed to obtain the current file directory")
	}

	envConfig := filepath.Dir(currentDirectory) + "/.env"
	err := godotenv.Load(envConfig)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDefaultPath() (path string) {
	path = utils.GetRunPath()
	path, _ = convert.GetString(utils.If(path != "", path, "/tmp"))
	return
}
