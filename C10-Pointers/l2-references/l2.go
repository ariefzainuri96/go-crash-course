package l2references

import (
	"strings"
)

func removeProfanity(message *string) {
	if strings.Contains(*message, "fubb") {
		*message = strings.ReplaceAll(*message, "fubb", "****")
	}

	if strings.Contains(*message, "shiz") {
		*message = strings.ReplaceAll(*message, "shiz", "****")
	}

	if strings.Contains(*message, "witch") {
		*message = strings.ReplaceAll(*message, "witch", "*****")
	}
}
