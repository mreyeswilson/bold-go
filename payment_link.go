package bold

import (
	"encoding/json"
	"log"
)

type AmountType string

const (
	AmountTypeOpen   AmountType = "OPEN"
	AmountTypeClosed AmountType = "CLOSE"
)

type TaxesOptions struct {
	Type  string  `json:"type"`
	Base  float64 `json:"base"`
	Value float64 `json:"value"`
}

type AmountTypeOptions struct {
	Currency    string        `json:"currency,omitempty"`
	Taxes       *TaxesOptions `json:"taxes,omitempty"`
	TipAmount   float64       `json:"tip_amount"`
	TotalAmount float64       `json:"total_amount"`
}

type PaymentLinkRequest struct {
	AmountType     AmountType         `json:"amount_type,omitempty"`
	Amount         *AmountTypeOptions `json:"amount,omitempty"`
	Description    string             `json:"description,omitempty"`
	ExpirationDate int64              `json:"expiration_date,omitempty"`
	CallbackUrl    string             `json:"callback_url,omitempty"`
	PaymentMethods []string           `json:"payment_methods,omitempty"`
	PayerEmail     string             `json:"payer_email,omitempty"`
	ImageUrl       string             `json:"image_url,omitempty"`
}

type PaymentLinkResponse struct {
	Payload struct {
		PaymentLink string `json:"payment_link"`
		Url         string `json:"url"`
	} `json:"payload"`
	Errors []any `json:"errors"`
}

func (b *Bold) GeneratePaymentLink(params *PaymentLinkRequest) (*PaymentLinkResponse, error) {
	data, err := b.doRequest("POST", "online/link/v1", params)
	if err != nil {
		log.Println("[Error] GeneratePaymentLink: ", err)
		return nil, err
	}

	paramObj, _ := json.Marshal(params)
	println(string(paramObj))

	var anydata map[string]any
	if err := json.Unmarshal(data, &anydata); err != nil {
		log.Println("[Error] Unmarshal: ", err)
		return nil, err
	}

	println(string(data))

	var response PaymentLinkResponse
	if err := json.Unmarshal(data, &response); err != nil {
		log.Println("[Error] Unmarshal: ", err)
		return nil, err
	}

	return &response, nil
}
