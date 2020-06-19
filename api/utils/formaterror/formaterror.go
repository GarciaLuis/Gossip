package formaterror

import (
	"errors"
	"fmt"
	"strings"
)

// FormatError helps format error messages into a readable form
func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname is already taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email is already in use")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title already exists")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	fmt.Println(err)

	return errors.New("Incorrect Details")
}
