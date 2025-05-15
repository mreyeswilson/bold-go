# Bold Go Client

Cliente en Go para interactuar con la API de Bold para procesamiento de pagos en Colombia.

## Instalación

```bash
go get github.com/mreyeswilson/bold-go
```

## Configuración

Primero, importa el paquete y crea una nueva instancia del cliente con tu API key:

```go
import "github.com/mreyeswilson/bold-go"

func main() {
    // Inicializar el cliente con tu API key de Bold
    client := bold.NewClient("tu-api-key-aquí")
    
    // Ahora puedes usar el cliente para hacer llamadas a la API
}
```

## Uso

### 1. Generar un enlace de pago

```go
// Crear una solicitud de enlace de pago
expirationDate, err := bold.FutureTimestampNs("24h")
if err != nil {
    log.Fatalf("Error al calcular la fecha de expiración: %v", err)
}

paymentRequest := &bold.PaymentLinkRequest{
    AmountType:  "CLOSE",
    Description: "Pago de productos en MiTienda.com",
    Amount: &bold.AmountTypeOptions{
        Currency:    "COP",
        TotalAmount: 100000.00, // $100,000 COP
        TipAmount:   0.00,
        Taxes: &bold.TaxesOptions{
            Type:  "IVA",
            Base:  84745.76,  // Base gravable
            Value: 15254.24,  // 19% de IVA en Colombia
        },
    },
    ExpirationDate: expirationDate, // Expira en 24 horas
    CallbackUrl:   "https://mitienda.com.co/webhook/payment-callback",
    PayerEmail:    "cliente@example.com.co",
    ImageUrl:      "https://mitienda.com.co/logo.png",
    PaymentMethods: []string{"credit_card", "debit_card", "pse"}, // Incluyendo PSE (Pagos Seguros en Línea)
}

// Generar el enlace de pago
response, err := client.GeneratePaymentLink(paymentRequest)
if err != nil {
    log.Fatalf("Error al generar enlace de pago: %v", err)
}

fmt.Printf("Enlace de pago generado: %s\n", response.Payload.Url)
```

### 2. Verificar el estado de un enlace de pago

```go
// Obtener el estado de un enlace de pago existente
status, err := client.GetPaymentLinkStatus("id-del-enlace")
if err != nil {
    log.Fatalf("Error al obtener estado del enlace: %v", err)
}

fmt.Printf("Estado del pago: %s\n", status.Status)
fmt.Printf("Monto total: $%.0f COP\n", status.Total) // Formato de moneda colombiana
fmt.Printf("Fecha de creación: %s\n", time.Unix(0, status.CreationDate).Format(time.RFC3339))

// Verificar si el pago fue exitoso
if status.Status == "PAID" {
    fmt.Println("¡Pago recibido exitosamente!")
}
```

### 3. Obtener métodos de pago disponibles

```go
// Obtener métodos de pago disponibles
methods, err := client.GetPaymentMethods()
if err != nil {
    log.Fatalf("Error al obtener métodos de pago: %v", err)
}

fmt.Println("Métodos de pago disponibles:")
for method, options := range methods.Payload.PaymentMethods {
    fmt.Printf("- %s (min: %d, max: %d)\n", method, options.Min, options.Max)
}
```

## Estructuras de Datos

### PaymentLinkRequest
Estructura para crear un nuevo enlace de pago:

```go
type PaymentLinkRequest struct {
    AmountType     string             `json:"amount_type,omitempty"`  // "OPEN" o "CLOSE"
    Amount         *AmountTypeOptions `json:"amount,omitempty"`
    Description    string             `json:"description,omitempty"`
    ExpirationDate int64              `json:"expiration_date,omitempty"` // Timestamp en nanosegundos
    CallbackUrl    string             `json:"callback_url,omitempty"`
    PaymentMethods []string           `json:"payment_methods,omitempty"`
    PayerEmail     string             `json:"payer_email,omitempty"`
    ImageUrl       string             `json:"image_url,omitempty"`
}
```

### PaymentLinkStatusResponse
Estructura con la respuesta del estado de un enlace de pago:

```go
type PaymentLinkStatusResponse struct {
    APIVersion     int32   `json:"api_version"`
    ID             string  `json:"id"`
    Total          float64 `json:"total"`
    Subtotal       float64 `json:"subtotal"`
    TipAmount      float64 `json:"tip_amount"`
    Taxes          []Tax   `json:"taxes"`
    Status         string  `json:"status"`  // ACTIVE, PROCESSING, PAID, REJECTED, CANCELLED, EXPIRED
    ExpirationDate int64   `json:"expiration_date"`
    CreationDate   int64   `json:"creation_date"`
    Description    string  `json:"description"`
    PaymentMethod  string  `json:"payment_method"`
    TransactionID  string  `json:"transaction_id"`
    AmountType     string  `json:"amount_type"`
    IsSandbox      bool    `json:"is_sandbox"`
}
```

## Consideraciones para Colombia

### Moneda
- La moneda por defecto es el Peso Colombiano (COP).
- Los montos se manejan sin decimales, siguiendo la convención local.

### Impuestos
- El IVA en Colombia es del 19% para la mayoría de bienes y servicios.
- Asegúrate de incluir correctamente la base gravable y el valor del IVA en las transacciones.

### Métodos de Pago
Los métodos de pago más comunes en Colombia incluyen:
- Tarjeta de crédito
- Tarjeta débito
- PSE (Pagos Seguros en Línea)
- Transferencia bancaria

## Manejo de Errores

Todas las funciones devuelven un error como segundo valor de retorno. Es importante verificar siempre este valor:

```go
response, err := client.GeneratePaymentLink(paymentRequest)
if err != nil {
    // Manejar el error
    log.Fatalf("Error: %v", err)
}
```

## Ambiente de Pruebas

Puedes utilizar el ambiente de sandbox de Bold para realizar pruebas. Simplemente usa una API key de prueba al inicializar el cliente.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, envía un pull request con tus cambios propuestos.

## Licencia

MIT