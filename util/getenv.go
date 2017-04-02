package util

import (
	"github.com/Sirupsen/logrus"
	"os"
)

func Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetenvNoDefault(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		logrus.Warn("Must provide env var: ", key)
	}
	return value
}
