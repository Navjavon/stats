package stats

import (
	"fmt"

	"github.com/Navjavon/bank/pkg/bank/types"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:     0,
			Amount: 10000,
			Category: "Avto",
		},
		{
			ID:     1,
			Amount: 10000,
			Category: "Restaurent",
		},
	}

	fmt.Println(TotalInCategory(payments, "Avto"))
	// Output: 10000
}