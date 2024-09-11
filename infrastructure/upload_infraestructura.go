package infrastructure

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file X_x.")
		return err
	}
	fmt.Println("Env loaded successfully O_o")
	return nil
}
