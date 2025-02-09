package main

import (
	"fmt"
	// c9 "main/CH9-Maps/c1-distict-words"
	utils "main/utils"

	mystrings "github.com/ariefzainuri96/mystrings"
)

type Car struct {
	color string
}

func (c *Car) setColor(color string) {
	c.color = color
}

func main() {
	// fmt.Println(CH3.Reformat("Hello World", CH3.Formater))
	// fizzbuzz()

	// testArray := [3]string{"test", "lower", "as"}
	// getMessageCosts(testArray[:])

	// checking()

	// showStrings("test", "lower")
	// showStrings(testArray[:]...)

	// isValidPassword("1234short")

	// c9.CountDistinctWords([]string{"Hello world", "hello there", "General Kenobi"})

	var x int = 50
	var y *int = &x
	*y = 100

	fmt.Println(x)
	fmt.Println(*y)

	c := Car{
		color: "red",
	}

	c.setColor("blue")

	fmt.Println(c.color)

	fmt.Println(mystrings.Reverse("Hello World"))
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}

func getMessageCosts(messages []string) {
	costs := make([]float64, 0, len(messages))

	for _, val := range messages {
		cost := float64(len(val)) * 0.01

		costs = append(costs, cost)
	}

	for _, value := range costs {
		fmt.Println(value)
	}
}

func checking() {
	cost := []string{}

	fmt.Println(len(cost))
	fmt.Println(cap(cost))
}

func showStrings(strings ...string) {
	for _, value := range strings {
		fmt.Println(value)
	}
}

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
		content := string(password[i])
		fmt.Printf("%v -> isDigit: %v, isUppercase: %v, isContainLetter: %v\n",
			content, utils.IsDigit(content), utils.IsUppercase(content), !utils.IsDigit(content))

		if !isContainLetter && !utils.IsDigit(content) {
			isContainLetter = true
		}

		if !isContainUppercase && utils.IsUppercase(content) {
			isContainUppercase = true
		}

		if !isContainDigit && utils.IsDigit(content) {
			isContainDigit = true
		}
	}

	return isContainLetter && isContainDigit && isContainUppercase
}
