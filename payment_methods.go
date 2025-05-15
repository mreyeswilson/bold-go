package bold

import "encoding/json"

type PaymentMethodOption struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type PaymentMethodResponse struct {
	Payload struct {
		PaymentMethods map[string]PaymentMethodOption `json:"payment_methods"`
		Errors         []any                          `json:"errors"`
	} `json:"payload"`
}

func (b *Bold) GetPaymentMethods() (*PaymentMethodResponse, error) {
	data, err := b.doRequest("GET", "online/link/v1/payment_methods", nil)
	if err != nil {
		return nil, err
	}

	var response PaymentMethodResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
