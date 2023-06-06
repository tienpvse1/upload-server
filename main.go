package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tienpvse1/upload-server/src/env"
	"github.com/tienpvse1/upload-server/src/modules/upload"
)

func ensureFolderExist(path string) {
	_, err := os.Open(path)
	if err != nil {
		os.MkdirAll(path, os.ModePerm)
	}
}

func main() {
	env.LoadEnv()
	ensureFolderExist(env.UPLOAD_DIR)
	app := fiber.New(fiber.Config{})

	app.Static(fmt.Sprintf("/%v", env.UPLOAD_DIR), env.UPLOAD_DIR)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Post(env.UPLOAD_ENDPOINT, upload.Upload)
	log.Fatal(app.Listen(env.PORT))
}
