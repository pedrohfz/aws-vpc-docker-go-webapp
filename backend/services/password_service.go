package services

import (
	"crypto/rand"
	"fmt"
)

const (
	passwordLength = 20
	lowerCase      = "abcdefghijklmnopqrstuvwxyz"
	upperCase      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
	symbols        = "!@#$%^&*()-_=+[]{};:,.?/"
)

func random(chars string) (byte, error) {
	randomByte := make([]byte, 1)
	if _, err := rand.Read(randomByte); err != nil {
		return 0, err
	}
	return chars[int(randomByte[0])%len(chars)], nil
}

func crypto(max int) (int, error) {
	randomByte := make([]byte, 1)
	if _, err := rand.Read(randomByte); err != nil {
		return 0, err
	}
	return int(randomByte[0]) % max, nil
}

func shuffle(b []byte) error {
	for i := len(b) - 1; i > 0; i-- {
		j, err := crypto(i + 1)
		if err != nil {
			return err
		}
		b[i], b[j] = b[j], b[i]
	}
	return nil
}

func GeneratePassword() (string, error) {
	charset := lowerCase + upperCase + digits + symbols
	passwordBytes := make([]byte, passwordLength)

	requiredCharSets := []string{lowerCase, upperCase, digits, symbols}

	writeIndex := 0

	for _, charSet := range requiredCharSets {
		ch, err := random(charSet)
		if err != nil {
			return "", fmt.Errorf("password generation failed: %w", err)
		}
		passwordBytes[writeIndex] = ch
		writeIndex++
	}

	for i := writeIndex; i < passwordLength; i++ {
		selectedChar, err := random(charset)
		if err != nil {
			return "", fmt.Errorf("password generation failed: %w", err)
		}
		passwordBytes[i] = selectedChar
	}

	if err := shuffle(passwordBytes); err != nil {
		return "", fmt.Errorf("password generation failed: %w", err)
	}

	return string(passwordBytes), nil
}
