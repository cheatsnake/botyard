package config

const (
	SqliteDatabasePath = "data/sqlite"
	SqliteDatabaseName = "store.db"
)

const (
	configFilename     = "./config/botyard.config.json"
	defaultPort        = "7007"
	defaultFilesFolder = "static"
)

var (
	defaultUserLimits = &userLimits{
		MinNicknameLength: 3,
		MaxNicknameLength: 32,
		AuthTokenLifetime: 10080,
	}

	defaultMessageLimits = &messageLimits{
		MaxBodyLength:    4096,
		MaxAttachedFiles: 10,
	}

	defaultFileLimits = &fileLimits{
		MaxImageSize: 2097152,
		MaxAudioSize: 5242880,
		MaxVideoSize: 26214400,
		MaxFileSize:  10485760,
	}
)
