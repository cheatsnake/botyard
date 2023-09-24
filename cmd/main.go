package main

import (
	"botyard/internal/config"
	"botyard/internal/http"
	"botyard/internal/logger"
	"botyard/internal/storage/memory"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	config.Load()
	config.LoadEnv()
	initDirsForFiles()

	appLogger := logger.New()
	storage := memory.New()
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
