package utils

import "log"

func MLogger(message string, statusCode int, err error) {
	log.Printf("{message: %s, code: %d, error: %s}\n", message, statusCode, err.Error())
}
