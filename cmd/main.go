package main

import (
	"log"
	"os"
	"path"

	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/http"
	"github.com/cheatsnake/botyard/internal/logger"
	"github.com/cheatsnake/botyard/internal/storage/sqlite"
)

func main() {
	appLogger := logger.New()
	err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	config.LoadEnv()
	config.InitDirsForFiles()
	config.InitDirsForSqlite()

	storage, err := sqlite.New(path.Join(config.SqliteDatabasePath, config.SqliteDatabaseName))
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
