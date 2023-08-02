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

	"github.com/gofiber/fiber/v2"
)

func main() {
	initDirsForFiles()

	app := fiber.New(fiber.Config{
		ErrorHandler: http.ErrHandler,
		BodyLimit:    25 * 1024 * 1024,
	})

	storage := memory.New()
	server := http.New(app, storage)
	server.App.Static("/", path.Join(".", "store"))
	server.InitRoutes("/api")

	go printMemoryUsage()

	log.Fatal(server.App.Listen(":4000"))
}

func bytesToMegabytes(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024
}

func printMemoryUsage() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		fmt.Printf("Current alloc: %.2f MB, total alloc: %.2f MB, GC cycles: %d\n",
			bytesToMegabytes(m.Alloc),
			bytesToMegabytes(m.TotalAlloc),
			m.NumGC,
		)
		time.Sleep(3 * time.Second)
	}
}

func initDirsForFiles() {
	fileDirName := "store"
	dirs := []string{"images", "videos", "audios", "files"}

	for _, dir := range dirs {
		err := os.MkdirAll(path.Join(".", fileDirName, dir), 0755)
		if err != nil {
			panic(fmt.Sprintf("Error creating directory: %s", err.Error()))
		}
	}
}
