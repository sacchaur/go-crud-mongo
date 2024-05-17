package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

// add function to encryp using sha256 the password based on the salt provided
func Encrypt(password, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))

	// return the hex encoded string
	return hex.EncodeToString(h.Sum(nil))
}

// add function to generate a random salt
func GenerateSalt() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(62)]
	}
	return string(b)
}

// add function to compare the password and the hash
func Compare(password, salt, hash string) bool {
	h := sha256.New()
	h.Write([]byte(password + salt))

	// return the comparison
	return hex.EncodeToString(h.Sum(nil)) == hash
}
