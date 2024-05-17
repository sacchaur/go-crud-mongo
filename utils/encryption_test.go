package utils

//Add test for the encryption.go file
import (
	"fmt"
	"testing"
)

// TestEncrypt function
func TestEncrypt(t *testing.T) {
	// create a new password
	password := "password"
	// create a new salt
	salt := GenerateSalt()
	// create a new hash
	hash := Encrypt(password, salt)
	// check if the hash is not empty
	fmt.Println(hash)
	if hash == "" {
		t.Error("Expected hash to not be empty")
	}
}

// TestGenerateSalt function
func TestGenerateSalt(t *testing.T) {
	// create a new salt
	salt := GenerateSalt()
	fmt.Println(salt)
	// check if the salt is not empty
	if salt == "" {
		t.Error("Expected salt to not be empty")
	}
}

// TestCompare function
func TestCompare(t *testing.T) {
	// create a new password
	password := "password"
	// create a new salt
	salt := GenerateSalt()
	// create a new hash
	hash := Encrypt(password, salt)
	// check if the hash is not empty
	if hash == "" {
		t.Error("Expected hash to not be empty")
	}
	// check if the password and hash match
	if !Compare(password, salt, hash) {
		t.Error("Expected password and hash to match")
	}
}
