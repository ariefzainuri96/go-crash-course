package c2passwordstrength

import (
	utils "main/utils"
)

/*
Password Strength

As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

    At least 5 characters long but no more than 12 characters.
    Contains at least one uppercase letter.
    Contains at least one digit.

A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the previous lesson to iterate over the characters of a string.

Assignment:

Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.
*/

func isValidPassword(password string) bool {
	const minChar = 5
	const maxChar = 12

	if len(password) < minChar || len(password) > maxChar {
		return false
	}

	isContainLetter := false
	isContainUppercase := false
	isContainDigit := false

	for i := 0; i < len(password); i++ {
		if !isContainLetter && !utils.IsDigit(string(password[i])) {
			isContainLetter = true
		}

		if !isContainUppercase && utils.IsUppercase(string(password[i])) {
			isContainUppercase = true
		}

		if !isContainDigit && utils.IsDigit(string(password[i])) {
			isContainDigit = true
		}
	}

	return isContainLetter && isContainDigit && isContainUppercase
}

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
