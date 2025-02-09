package l6nilpointers

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
