package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func GetEnv(key string, filename ...string) string {
	err := godotenv.Load(filename...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	value := os.Getenv(key)
	return value
}

func GetTimeNow() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func S2json(value any) string {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("%s %v", GetTimeNow(), err)
	}
	var str bytes.Buffer
	err = json.Indent(&str, jsonBytes, "", "\t")
	if err != nil {
		fmt.Printf("%s %v", GetTimeNow(), err)
	}
	result := str.String()
	return result
}
