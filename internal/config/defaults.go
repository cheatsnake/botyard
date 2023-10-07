package config

const (
	SqliteDatabasePath = "data/sqlite"
	SqliteDatabaseName = "store.db"
)

const (
	configFilename     = "botyard.config.json"
	defaultPort        = "4000"
	defaultFilesFolder = "static"
)

var (
	defaultUserLimits = &userLimits{
		MinNicknameLength: 3,
		MaxNicknameLength: 32,
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
