package gateway_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/gateway"
)

func TestFinishOrder(t *testing.T) {
	orderID := "12345"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := fmt.Sprintf("/order/%s/status/COMPLETED", orderID)
		if r.URL.Path != expectedURL {
			t.Errorf("unexpected URL: got %s, want %s", r.URL.Path, expectedURL)
		}

		if r.Method != "PUT" {
			t.Errorf("unexpected request method: got %s, want PUT", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	orderApi := gateway.NewOrderApi(gateway.OrderApiConfig{
		OrderApiBaseUrl: server.URL,
	})

	err := orderApi.FinishOrder(orderID)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestReleaseOrder(t *testing.T) {
	orderID := "12345"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := fmt.Sprintf("/order/%s/status/READY", orderID)
		if r.URL.Path != expectedURL {
			t.Errorf("unexpected URL: got %s, want %s", r.URL.Path, expectedURL)
		}

		if r.Method != "PUT" {
			t.Errorf("unexpected request method: got %s, want PUT", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	orderApi := gateway.NewOrderApi(gateway.OrderApiConfig{
		OrderApiBaseUrl: server.URL,
	})

	err := orderApi.ReleaseOrder(orderID)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
