package main

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	ErrUsernameTooShort          = errors.New("Username is too short!")
	ErrUsernameEmpty             = errors.New("Username cannot be empty!")
	ErrUsernameInvalidCharacters = errors.New("Invalid characters in the username!")
)

type UsernameValidationError struct {
	Username string
	Reason   string
	Err      error
}

// Unwrap the err field from the usernamevalidationError struct
func (e *UsernameValidationError) Unwrap() error {
	return e.Err
}

func (e *UsernameValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Username validation error for %v: %v", e.Username, e.Err)
	}
	return fmt.Sprintf("Username error validation error for %v, Reason: %v, Details(%v)", e.Username, e.Reason, e.Err)
}

func ValidateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("Username validation error: %w", ErrUsernameEmpty)
	}
	if len(username) < 3 {
		return fmt.Errorf("Username validation error for %v: %w", username, ErrUsernameTooShort)
	}

	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return &UsernameValidationError{
				Username: username,
				Reason:   "contains non-alphanumeric character",
				Err:      fmt.Errorf("invalid character: %v, %w", string(r), ErrUsernameInvalidCharacters),
			}
		}
	}
	return nil
}

func RegisterUser(username, password string) (string, error) {
	if err := ValidateUsername(username); err != nil {
		return "", fmt.Errorf("User registration failed for %v, %w", username, err)
	}
	if len(password) < 6 {
		return "", fmt.Errorf("Password is too short!")
	}
	return "You have been registered successfully!", nil
}

func main() {
	// Scenerio 1: Username validation
	fmt.Println("-----Scenerio 1: Username validation-----")
	msg, err := RegisterUser("OsitadinmaDivine!", "123")
	if err != nil {
		fmt.Println("Registration Error: ", err)
		var uvErr *UsernameValidationError
		if errors.Is(err, ErrUsernameEmpty) {
			fmt.Println("Specific error for empty username!")
		}
		if errors.Is(err, ErrUsernameTooShort) {
			fmt.Println("Specific error for short username!")
		}
		if errors.As(err, &uvErr) {
			fmt.Println("Specific error for a username that invalid characters!")
		}
	} else {
		fmt.Println(msg)
	}
}
