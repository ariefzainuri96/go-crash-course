package c1

func countConnections(groupSize int) int {
	numConnect := 0

	for i := 1; i < groupSize; i++ {
		numConnect += groupSize - i
	}

	return numConnect
}
