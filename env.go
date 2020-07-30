package aefire

import (
	"os"
	"strconv"
)

func GetEnvInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		value = defaultValue
	}

	return value
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	} else {
		return defaultValue
	}
}
