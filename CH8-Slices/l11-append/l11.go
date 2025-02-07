package l11append

/*
We’ve been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day’s costs.

If there are no costs for that day, return an empty slice.
*/

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	if len(costs) == 0 {
		return []float64{}
	}

	daysCost := make([]float64, 0, len(costs))

	for _, value := range costs {
		if day != value.day {
			continue
		}

		daysCost = append(daysCost, value.value)
	}

	return daysCost
}
