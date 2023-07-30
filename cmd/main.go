package main

import (
	"botyard/internal/http"
	"botyard/internal/storage/memory"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: http.ErrHandler,
	})

	storage := memory.New()
	server := http.New(app, storage)
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
