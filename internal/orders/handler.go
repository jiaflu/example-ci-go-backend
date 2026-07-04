package orders

import (
	"encoding/json"
	"net/http"
)

func HandleSummary(w http.ResponseWriter, r *http.Request) {
	summary, err := Summarize([]Order{
		{ID: "ord-1001", Amount: 1200, Paid: true},
		{ID: "ord-1002", Amount: 900, Paid: false},
		{ID: "ord-1003", Amount: 2500, Paid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(summary)
}
