package c1distictwords

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]bool)

	for _, value := range messages {

		splittedValue := strings.Fields(value)

		for _, sValue := range splittedValue {
			_, isExist := distinctWords[strings.ToLower(sValue)]

			if !isExist {
				distinctWords[strings.ToLower(sValue)] = true
			}
		}
	}

	return len(distinctWords)
}

func CountDistinctWords(messages []string) int {
	distinctWords := make(map[string]bool)

	for _, value := range messages {

		splittedValue := strings.Fields(value)

		fmt.Println(splittedValue)

		for _, sValue := range splittedValue {
			fmt.Println(sValue)

			_, isExist := distinctWords[strings.ToLower(sValue)]

			if !isExist {
				distinctWords[strings.ToLower(sValue)] = true
			}
		}
	}

	return len(distinctWords)
}
