package billing

import (
	"errors"
)

// OpcionPago simula diferentes métodos de pago.
type OpcionPago string

const (
	TarjetaCredito OpcionPago = "Tarjeta de Crédito"
	PayPal         OpcionPago = "PayPal"
)

// Pagable define la interfaz para cualquier proceso de pago.
type Pagable interface {
	RegistrarPago(monto float32, opcion OpcionPago, suscripcion *Suscripcion) error
	VerificarPago(pagoID string) bool
}

// ErrPagoFallido es un error para cuando el proceso de pago no es exitoso.
var ErrPagoFallido = errors.New("el pago ha fallado o fue rechazado")

// Implementación de la interfaz Pagable en Pago
// ... (Ver archivo pago.go)
