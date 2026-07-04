package orders

import "errors"

var ErrNoOrders = errors.New("no orders provided")

type Order struct {
	ID     string
	Amount int
	Paid   bool
}

type Summary struct {
	TotalOrders int `json:"totalOrders"`
	PaidOrders  int `json:"paidOrders"`
	Revenue     int `json:"revenue"`
}

func Summarize(items []Order) (Summary, error) {
	if len(items) == 0 {
		return Summary{}, ErrNoOrders
	}

	var summary Summary
	for _, item := range items {
		summary.TotalOrders++
		if item.Paid {
			summary.PaidOrders++
			summary.Revenue += item.Amount
		}
	}

	return summary, nil
}
