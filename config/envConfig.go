package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config retrieves the value of a given environment variable.
//
// Parameters: key (string) - the name of the environment variable, defaultValue (string) - the default value to return if the variable is not set.
// Returns: string - the value of the environment variable or the default value if it's not set.
func Config(key string, defaultValue string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
