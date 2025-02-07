package l2nestedstructsingo

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// sender
	if mToSend.sender.number == 0 || mToSend.sender.name == "" {
		return false
	}

	// recipient
	if mToSend.recipient.number == 0 || mToSend.recipient.name == "" {
		return false
	}

	return true
}
