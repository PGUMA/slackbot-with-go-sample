package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	fmt.Println(os.Getenv("SLACK_SIGNING_SECRET"))
}

// .envを呼び出します。
func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("can not load env file: %v", err)
	}
}
