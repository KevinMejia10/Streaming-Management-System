package billing

import (
	"fmt"
	"time"
)

// Pago representa una transacción de pago dentro del sistema.
type Pago struct {
	id         string
	monto      float32
	opcion     OpcionPago
	verificado bool
}

// NuevoPago crea una nueva instancia de Pago.
func NuevoPago(id string, monto float32, opcion OpcionPago) *Pago {
	return &Pago{
		id:         id,
		monto:      monto,
		opcion:     opcion,
		verificado: false,
	}
}

// RegistrarPago simula el procesamiento y registro de un pago.
func (p *Pago) RegistrarPago(monto float32, opcion OpcionPago, suscripcion *Suscripcion) error {
	// Simulación de interacción con pasarela de pago
	fmt.Printf("💳 Procesando pago de $%.2f con %s...\n", monto, opcion)

	if monto < 0 {
		return ErrPagoFallido
	}

	p.monto = monto
	p.opcion = opcion
	p.id = time.Now().Format("060102150405") // Simulación de ID de transacción

	// Simular éxito del pago
	p.verificado = true

	// Lógica de negocio: actualizar la suscripción
	suscripcion.Renovar() // Renovar activa y extiende la fecha fin

	fmt.Printf("✅ Pago ID %s registrado con éxito. Suscripción actualizada a Activa hasta %s.\n", p.id, suscripcion.fechaFin.Format("02-Jan-2006"))
	return nil
}

// VerificarPago simula la consulta del estado de un pago.
func (p *Pago) VerificarPago(pagoID string) bool {
	return p.verificado && p.id == pagoID
}
