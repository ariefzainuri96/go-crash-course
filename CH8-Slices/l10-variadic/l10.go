package l10variadic

func sum(nums ...int) int {
	cost := 0

	for _, value := range nums {
		cost += value
	}

	return cost
}
