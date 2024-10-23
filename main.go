package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func loadEnv() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func main(){
    loadEnv()
    fmt.Println("Everything working fine.!!!")
}