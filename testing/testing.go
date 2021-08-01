package testing

import (
	"math/rand"
	"time"
)

const (
	chars    = "abcdefghijklmnopqrstuvwxyz0123456789"
	Protocol = "http"
	Host     = "localhost"
	Port     = "4002"
	// Port     = "4000"
)

func GetRandomString() string {
	length := 6
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func GetRandomEmail() string {
	return GetRandomString() + "@localhost.com"
}
