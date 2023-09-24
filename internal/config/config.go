package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Service service `json:"service"`

	Limits struct {
		User    userLimits    `json:"user,omitempty"`
		Message messageLimits `json:"message,omitempty"`
		File    fileLimits    `json:"file,omitempty"`
	} `json:"limits,omitempty"`
}

type service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar,omitempty"`

	Socials []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"socials"`
}

type userLimits struct {
	MinNicknameLength int `json:"minNicknameLength,omitempty"`
	MaxNicknameLength int `json:"maxNicknameLength,omitempty"`
}

type messageLimits struct {
	MaxBodyLength    int `json:"maxBodyLength,omitempty"`
	MaxAttachedFiles int `json:"maxAttachedFiles,omitempty"`
}

type fileLimits struct {
	MaxImageSize int `json:"maxImageSize,omitempty"`
	MaxAudioSize int `json:"maxAudioSize,omitempty"`
	MaxVideoSize int `json:"maxVideoSize,omitempty"`
	MaxFileSize  int `json:"maxFileSize,omitempty"`
}

var GlobalConfig Config

// Load main app config
func Load() {
	var conf Config

	confFile, err := os.Open(configFilename)
	if err != nil {
		panic(fmt.Sprintf("failed reading config file: %s", err.Error()))
	}

	defer confFile.Close()

	jsonParser := json.NewDecoder(confFile)
	if err := jsonParser.Decode(&conf); err != nil {
		panic(fmt.Sprintf("failed parsing config file: %s", err.Error()))
	}

	if conf.Limits.User.MinNicknameLength == 0 {
		conf.Limits.User.MinNicknameLength = defaultUserLimits.MinNicknameLength
	}

	if conf.Limits.User.MaxNicknameLength == 0 {
		conf.Limits.User.MaxNicknameLength = defaultUserLimits.MaxNicknameLength
	}

	if conf.Limits.Message.MaxBodyLength == 0 {
		conf.Limits.Message.MaxBodyLength = defaultMessageLimits.MaxBodyLength
	}

	if conf.Limits.Message.MaxAttachedFiles == 0 {
		conf.Limits.Message.MaxAttachedFiles = defaultMessageLimits.MaxAttachedFiles
	}

	if conf.Limits.File.MaxImageSize == 0 {
		conf.Limits.File.MaxImageSize = defaultFileLimits.MaxImageSize
	}

	if conf.Limits.File.MaxAudioSize == 0 {
		conf.Limits.File.MaxAudioSize = defaultFileLimits.MaxAudioSize
	}

	if conf.Limits.File.MaxVideoSize == 0 {
		conf.Limits.File.MaxVideoSize = defaultFileLimits.MaxVideoSize
	}

	if conf.Limits.File.MaxFileSize == 0 {
		conf.Limits.File.MaxFileSize = defaultFileLimits.MaxFileSize
	}

	GlobalConfig = conf
}

// Load environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed parsing .env file: %s", err.Error()))
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", defaultPort)
	}

	if os.Getenv("FILES_FOLDER") == "" {
		os.Setenv("FILES_FOLDER", defaultFilesFolder)
	}
}
