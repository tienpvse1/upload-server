package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT            string = ""
	UPLOAD_DIR      string = ""
	UPLOAD_ENDPOINT string = ""
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load environment file")
	}
	PORT = fmt.Sprintf(":%v", os.Getenv("GO_PORT"))
	UPLOAD_DIR = os.Getenv("GO_UPLOAD_DIR")
	UPLOAD_ENDPOINT = os.Getenv("GO_UPLOAD_ENDPOINT")
	fmt.Printf("upload directory: %v \n", UPLOAD_DIR)
	fmt.Printf("upload endpoint: %v \n", UPLOAD_ENDPOINT)
}
