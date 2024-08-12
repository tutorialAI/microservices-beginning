package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	githubUsername, exists := os.LookupEnv("GITHUB_USERNAME")

	if exists {
		fmt.Println(githubUsername)
	}

	githubAPIKey, exists := os.LookupEnv("GITHUB_API_KEY")

	if exists {
		fmt.Println(githubAPIKey)
	}
}
