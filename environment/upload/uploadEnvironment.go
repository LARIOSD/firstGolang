package upload

import (
	"fmt"
	"github.com/joho/godotenv"
)

func UploadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Could not load .env file X_x.\n")
	}
	fmt.Println("Env loaded successfully O_o")
}
