package stats

import "github.com/Navjavon/bank/pkg/bank/types"

// TotalInCategory finds the amount of purchases in a specific category.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}

	return sum
}
