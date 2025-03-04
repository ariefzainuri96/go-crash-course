package main

import (
	"fmt"
	// c9 "main/CH9-Maps/c1-distict-words"
	utils "main/utils"

	api "main/services"

	// l7select "main/CH12-Channels/l7-select"

	ch12c1 "main/CH12-Channels/c1"
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

	// var x int = 50
	// var y *int = &x
	// *y = 100

	// fmt.Println(x)
	// fmt.Println(*y)

	// c := Car{
	// 	color: "red",
	// }

	// c.setColor("blue")

	// fmt.Println(c.color)

	// fmt.Println(mystrings.Reverse("Hello World"))

	// c12l1.ConcurrencyTest("Hello World 1")
	// c12l1.ConcurrencyTest("Hello World 2")
	// c12l1.ConcurrencyTest("Hello World 3")
	// getUserData()

	// testAddChannel()

	// l7select.Test(
	// 	[]string{
	// 		"hi friend",
	// 		"What's going on?",
	// 		"Welcome to the business",
	// 		"I'll pay you to be my friend",
	// 	},
	// 	[]string{
	// 		"Will you make your appointment?",
	// 		"Let's be friends",
	// 		"What are you doing?",
	// 		"I can't believe you've done this.",
	// 	},
	// )
	// l7select.Test(
	// 	[]string{
	// 		"this song slaps hard",
	// 		"yooo hoooo",
	// 		"i'm a big fan",
	// 	},
	// 	[]string{
	// 		"What do you think of this song?",
	// 		"I hate this band",
	// 		"Can you believe this song?",
	// 	},
	// )

	ch12c1.Test(4)
}

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	for _, email := range emails {
		ch <- email
	}

	close(ch)

	return ch
}

func testAddChannel() {
	emails := []string{
		"To boldly go where no man has gone before.",
		"Live long and prosper.",
	}

	ch := addEmailsToQueue(emails)

	for email := range ch {
		fmt.Println(email)
	}
}

func getUserData() {
	userCh := make(chan api.User)
	errCh := make(chan error)

	fmt.Println("Fetching User Data...")

	go api.GetUserData(userCh, errCh)

	select {
	case err := <-errCh:
		fmt.Println("Error:", err)
	case user := <-userCh:
		fmt.Println("User:", user)
	}
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
