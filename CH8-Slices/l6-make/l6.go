package l6make

func getMessageCosts(messages []string) []float64 {
	// here we set default empty slices with capacity of messages length
	// it is usefull for memory management
	costs := make([]float64, 0, len(messages))

	for _, val := range messages {
		cost := float64(len(val)) * 0.01

		costs = append(costs, cost)
	}

	return costs
}
