package utils

import (
	"fmt"
	"os"
)

func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("key `%v` is not set in environment variables", key))
	}

	return val

}

func Get(key, default_value string) string {

	val := os.Getenv(key)
	if val == "" {
		return default_value
	}

	return val

}
