package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/http"
	"github.com/cheatsnake/botyard/internal/logger"
	"github.com/cheatsnake/botyard/internal/storage/sqlite"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	config.LoadEnv()
	initDirsForFiles()

	appLogger := logger.New()
	storage, err := sqlite.New("data/sqlite/store.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.InitTables(); err != nil {
		log.Fatal(err)
	}

	server := http.New(storage)
	server.InitRoutes()

	go appLogger.MemoryUsage()
	log.Fatal(server.App.Listen(":" + os.Getenv("PORT")))
}

func initDirsForFiles() {
	dirs := []string{"images", "videos", "audios", "files"}

	for _, dir := range dirs {
		err := os.MkdirAll(path.Join(".", os.Getenv("FILES_FOLDER"), dir), 0755)
		if err != nil {
			panic(fmt.Sprintf("failed to create a directory for files: %s", err.Error()))
		}
	}
}
