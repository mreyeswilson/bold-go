package bold

import "time"

// FutureTimestampNs genera un timestamp en nanosegundos sumando la duración especificada al tiempo actual.
// La duración debe ser un string válido según time.ParseDuration (ej: "24h", "30m", "48h").
// Ejemplo de uso:
//   timestamp, err := FutureTimestampNs("24h") // Expira en 24 horas
func FutureTimestampNs(durationStr string) (int64, error) {
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        return 0, err
    }
    futureTime := time.Now().Add(duration)
    return futureTime.UnixNano(), nil
}
