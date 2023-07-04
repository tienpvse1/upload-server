package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tienpvse1/upload-server/src/env"
)

func asynchronousSave(channel chan UploadInfo, c *fiber.Ctx, file *multipart.FileHeader) {
	filename := fmt.Sprintf("%v_%v", uuid.New().String(), file.Filename)
	err := c.SaveFile(file, fmt.Sprintf("%v/%v", env.UPLOAD_DIR, filename))
	if err != nil {
		log.Fatal(err)
		channel <- UploadInfo{
			Filename: filename,
			Path:     "Cannot upload this file",
		}
		return
	}
	channel <- UploadInfo{
		Filename: filename,
		Path:     fmt.Sprintf("%v/%v", env.UPLOAD_DIR, filename),
	}

}

func Upload(c *fiber.Ctx) error {
	result := []UploadInfo{}
	channel := make(chan UploadInfo)
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	for _, file := range files {
		go asynchronousSave(channel, c, file)
	}
	for _, file := range files {
		result = append(result, <-channel)
		fmt.Println(file.Filename)
	}
	return c.JSON(UploadResponse{
		Success: true,
		Status:  201,
		Result: UploadResult{
			Success: true,
			Result:  result,
		},
		Errors:    nil,
		Timestamp: time.Now().UTC(),
	})
}
