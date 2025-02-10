package l4bufferedchannels

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	for _, email := range emails {
		ch <- email
	}

	// close the channel when the emails loop is ended
	// this will also tell the consumers that the channel is closed
	// This is critical because it tells the consumer that no more values will come, allowing the process of streaming data from channel is exit gracefully.
	close(ch)

	return ch
}
