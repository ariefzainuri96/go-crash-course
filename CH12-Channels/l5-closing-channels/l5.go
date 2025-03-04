package l5closingchannels

func countReports(numSentCh chan int) int {
	totalNumSent := 0

	for value := range numSentCh {
		totalNumSent += value
	}

	return totalNumSent
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
