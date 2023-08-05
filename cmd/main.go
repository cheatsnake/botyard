package main

import (
	"botyard/internal/http"
	"botyard/internal/storage/memory"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

const filesFolder = "stock"
const host = ""
const port = "4000"

func main() {
	initEnv()
	initDirsForFiles()

	storage := memory.New()
	server := http.New(storage)
	server.App.Static("/"+filesFolder, path.Join(".", filesFolder))
	server.InitRoutes("/api")

	go printMemoryUsage()

	log.Fatal(server.App.Listen(host + ":" + port))
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}
}

func initDirsForFiles() {
	dirs := []string{"images", "videos", "audios", "files"}

	for _, dir := range dirs {
		err := os.MkdirAll(path.Join(".", filesFolder, dir), 0755)
		if err != nil {
			panic(fmt.Sprintf("Error creating directory: %s", err.Error()))
		}
	}
}

func printMemoryUsage() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		fmt.Printf("Current alloc: %.2f MB, total alloc: %.2f MB, GC cycles: %d\n",
			float64(m.Alloc)/1024/1024,
			float64(m.TotalAlloc)/1024/1024,
			m.NumGC,
		)
		time.Sleep(3 * time.Second)
	}
}
