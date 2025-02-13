package l3passbyreference

/*

Assignment

Write a getMessageText function. It should accept a pointer to an Analytics struct and a Message struct (not a pointer).

It should look at the Success field of the Message struct and, based on that, increment the MessagesTotal, MessagesSucceeded, or MessagesFailed fields of the Analytics struct as appropriate.

*/

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytics *Analytics, message Message) {
	(*analytics).MessagesTotal++

	if message.Success {
		(*analytics).MessagesSucceeded++
	} else {
		(*analytics).MessagesFailed++
	}
}
