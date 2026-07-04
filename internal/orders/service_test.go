package orders

import (
	"errors"
	"testing"
)

func TestSummarize(t *testing.T) {
	summary, err := Summarize([]Order{
		{ID: "ord-1", Amount: 100, Paid: true},
		{ID: "ord-2", Amount: 50, Paid: false},
		{ID: "ord-3", Amount: 75, Paid: true},
	})
	if err != nil {
		t.Fatalf("summarize returned error: %v", err)
	}

	if summary.TotalOrders != 3 {
		t.Fatalf("expected 3 total orders, got %d", summary.TotalOrders)
	}

	if summary.PaidOrders != 2 {
		t.Fatalf("expected 2 paid orders, got %d", summary.PaidOrders)
	}

	if summary.Revenue != 175 {
		t.Fatalf("expected revenue 175, got %d", summary.Revenue)
	}
}

func TestSummarizeRejectsEmptyInput(t *testing.T) {
	_, err := Summarize(nil)
	if !errors.Is(err, ErrNoOrders) {
		t.Fatalf("expected ErrNoOrders, got %v", err)
	}
}
