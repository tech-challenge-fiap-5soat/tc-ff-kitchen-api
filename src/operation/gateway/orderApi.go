package gateway

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
)

type orderApi struct {
	httpClient      http.Client
	orderApiBaseURL string
}

type OrderApiConfig struct {
	OrderApiBaseUrl string
}

func NewOrderApi(config OrderApiConfig) interfaces.OrderApi {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &orderApi{
		httpClient:      *client,
		orderApiBaseURL: config.OrderApiBaseUrl,
	}
}

func (oa *orderApi) ReleaseOrder(orderId string) error {
	return oa.updateOrder(orderId, "READY")
}

func (oa *orderApi) FinishOrder(orderId string) error {
	return oa.updateOrder(orderId, "COMPLETED")
}

func (oa *orderApi) updateOrder(orderId string, status string) error {
	url := fmt.Sprintf("%s/order/%s/status/%s", oa.orderApiBaseURL, orderId, status)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return fmt.Errorf("error occurred while creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := oa.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error occurred while sending put request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: got %v", resp.StatusCode)
	}

	return nil
}
