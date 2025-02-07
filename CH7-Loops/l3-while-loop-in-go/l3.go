package l3whileloopingo

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1

	balance := float64(maxCostInPennies) - actualCostInPennies

	for balance >= 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}

	maxMessagesToSend--

	return maxMessagesToSend
}
