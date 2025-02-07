package ch3functions

import (
	"fmt"
)

func Reformat(message string, formatter func(string) string) string {
	return fmt.Sprintf("TEXTIO: %s", formatter(message))
}

func Formater(message string) string {
	return fmt.Sprintf("%s...", message)
}

// when you start the function with lowercase, this will not be visible when we try to access from another package
// but it is still available in same package
func testingLowercase() {
	fmt.Println("this is testingLowercase")
}
