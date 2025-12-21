package billing

import "errors"

// OpcionPago define las opciones de pago.
type OpcionPago string

const (
	TarjetaCredito OpcionPago = "Tarjeta"
	PayPal         OpcionPago = "PayPal"
)

// ErrPagoFallido es un error para representar transacciones fallidas.
var ErrPagoFallido = errors.New("el pago ha fallado o es inválido")
