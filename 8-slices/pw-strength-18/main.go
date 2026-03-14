// Textio wants to enforce a new password policy. A valid password must meet the following criteria:

//     At least 5 characters long but no more than 12 characters.
//     Contains at least one uppercase letter.
//     Contains at least one digit.

// Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

package main

import (
	"unicode"
)

func isValidPassword(password string) bool {
	// criteria 1 -- why do they want me to use loop for this lmao
	length_criteria := !(len(password) >= 5 && len(password) < 12)
	if length_criteria {
		return false
	}

	// even these 2 could be checked with Slice contains ig
	// criteria 2 - upper case
	uppercase_criteria := false
	// criteria 3 - one digit
	digit_criteria := false

	for _, pw := range password {
		switch {
		case unicode.IsUpper(pw):
			uppercase_criteria = true

		case unicode.IsDigit(pw):
			digit_criteria = true
		}
	}

	if uppercase_criteria && digit_criteria {
		return true
	}

	return false
}
