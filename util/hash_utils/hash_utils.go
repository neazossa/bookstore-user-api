package hash_utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash password using the Bcrypt hashing algorithm
// and then return the hashed password as a
// base64 encoded string
func HashPassword(password string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return string(hashedPasswordBytes)
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
