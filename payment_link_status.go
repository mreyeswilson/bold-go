package bold

import "encoding/json"

type Tax struct {
	Type  string  `json:"type"`  // Ej: "VAT"
	Base  float64 `json:"base"`  // Monto base del impuesto
	Value float64 `json:"value"` // Valor final del impuesto
}

type PaymentLinkStatusResponse struct {
	APIVersion     int32   `json:"api_version"`     // Ej: "1.0"
	ID             string  `json:"id"`              // Ej: "123456789"
	Total          float64 `json:"total"`           // Ej: 150.00
	Subtotal       float64 `json:"subtotal"`        // Ej: 130.00
	TipAmount      float64 `json:"tip_amount"`      // Ej: 5.00
	Taxes          []Tax   `json:"taxes"`           // Lista de impuestos
	Status         string  `json:"status"`          // ACTIVE, PROCESSING, PAID, REJECTED, CANCELLED, EXPIRED
	ExpirationDate int64   `json:"expiration_date"` // UNIX timestamp. Ej: 1627846267000
	CreationDate   int64   `json:"creation_date"`   // UNIX timestamp. Ej: 1719242727607215713
	Description    string  `json:"description"`     // Ej: "Compra de productos electrónicos"
	PaymentMethod  string  `json:"payment_method"`  // Ej: "credit_card"
	TransactionID  string  `json:"transaction_id"`  // Ej: "txn_001"
	AmountType     string  `json:"amount_type"`     // OPEN o CLOSE
	IsSandbox      bool    `json:"is_sandbox"`      // true si es entorno de pruebas
}

func (b *Bold) GetPaymentLinkStatus(id string) (*PaymentLinkStatusResponse, error) {
	data, err := b.doRequest("GET", "online/link/v1/"+id, nil)
	if err != nil {
		return nil, err
	}

	var anydata map[string]any
	if err := json.Unmarshal(data, &anydata); err != nil {
		return nil, err
	}

	println(string(data))

	var response PaymentLinkStatusResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
