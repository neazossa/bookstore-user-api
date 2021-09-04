package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash password using the Bcrypt hashing algorithm
// and then return the hashed password as a
// base64 encoded string
func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func main() {
	// Hash password
	var hashedPassword, err = hashPassword("password1")

	if err != nil {
		println(fmt.Println("Error hashing password"))
		return
	}

	fmt.Println("Password Hash:", hashedPassword)

	// Check if passed password matches the original password
	fmt.Println("Password Match:", doPasswordsMatch(hashedPassword, "password1"))
}
