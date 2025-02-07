package l5countinstances

/*
Assignment

Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the getCounts function. It takes a slice of strings messagedUsers and a map of string -> int validUsers. It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if “benji” is in the map and appears in the slice 3 times, the key “benji” in the map should have the value 3.
*/

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, value := range messagedUsers {

		if _, isExist := validUsers[value]; isExist {
			validUsers[value] += 1
		}

	}
}
