package l1loopsingo

import (
	"fmt"
	"strconv"
)

func bulkSend(numMessages int) float64 {
	cost := 1.0 * float64(numMessages)
	additionalCostPerMessage := 0.0

	for i := 0; i < numMessages; i++ {
		additionalCostPerMessage += 0.01 * float64(i)
	}

	formattedCost, error := strconv.ParseFloat(fmt.Sprintf("%.2f", cost+additionalCostPerMessage), 64)

	if error != nil {
		return cost + additionalCostPerMessage
	}

	return formattedCost
}
