package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars     = lowerChars + upperChars + digits + specialChars
)

func GenerateStrongResetCode(length int) (string, error) {
	if length < 8 {
		return "", fmt.Errorf("password length should be at least 8 characters")
	}

	password := make([]byte, length)

	categories := []string{lowerChars, upperChars, digits, specialChars}
	for i := 0; i < len(categories); i++ {
		char, err := randomChar(categories[i])
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	for i := len(categories); i < length; i++ {
		char, err := randomChar(allChars)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	shuffle(password)

	return string(password), nil
}

func randomChar(chars string) (byte, error) {
	max := big.NewInt(int64(len(chars)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return chars[n.Int64()], nil
}

func shuffle(data []byte) {
	for i := len(data) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			continue
		}
		data[i], data[j.Int64()] = data[j.Int64()], data[i]
	}
}
