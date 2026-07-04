package orders

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleSummary(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/orders/summary", nil)
	rec := httptest.NewRecorder()

	HandleSummary(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var body Summary
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Revenue != 3700 {
		t.Fatalf("expected revenue 3700, got %d", body.Revenue)
	}
}
