package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	sRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		index := sRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}
