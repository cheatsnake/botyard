package config

import (
	"fmt"
	"os"
	"path"
)

func InitDirsForFiles() {
	dirs := []string{"images", "videos", "audios", "files"}

	for _, dir := range dirs {
		err := os.MkdirAll(path.Join(".", os.Getenv("FILES_FOLDER"), dir), 0755)
		if err != nil {
			panic(fmt.Sprintf("failed to create a directory for files: %s", err.Error()))
		}
	}
}

func InitDirsForSqlite() {
	err := os.MkdirAll(SqliteDatabasePath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("failed to create a directory for database: %s", err.Error()))
	}
}
