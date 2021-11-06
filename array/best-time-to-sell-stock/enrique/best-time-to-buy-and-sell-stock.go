package enrique

func maxProfit(prices []int) int {
	cheapestSoFar := prices[0]
	profitSoFar := 0

	for _, p := range prices {
		if p < cheapestSoFar {
			cheapestSoFar = p
			continue
		}

		possibleProfit := p - cheapestSoFar
		if possibleProfit > profitSoFar {
			profitSoFar = possibleProfit
		}

	}

	return profitSoFar
}