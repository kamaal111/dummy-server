package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func GetBytes(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func MLogger(message string, statusCode int, err error) {
	log.Printf("{message: %s, code: %d, error: %s}\n", message, statusCode, err.Error())
}
