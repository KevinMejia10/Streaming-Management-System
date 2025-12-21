package billing

import (
	"fmt"
	"time"
)

// Pago representa una transacción de pago dentro del sistema.
type Pago struct {
	id        string
	monto     float64
	opcion    OpcionPago // Usa la definición de billing.go
	fechaPago time.Time
	estado    string // PENDIENTE, COMPLETADO, FALLIDO
}

// NuevoPago crea una nueva instancia de Pago.
func NuevoPago(id string, monto float64, opcion OpcionPago) *Pago {
	return &Pago{
		id:     id,
		monto:  monto,
		opcion: opcion,
		estado: "PENDIENTE",
	}
}

// RegistrarPago simula el procesamiento y registro de un pago.
func (p *Pago) RegistrarPago(monto float64, opcion OpcionPago, suscripcion *Suscripcion) error {

	if monto <= 0 {
		p.estado = "FALLIDO"
		return ErrPagoFallido // Usa el error definido en billing.go
	}

	fmt.Printf("💳 Procesando pago de $%.2f con %s...\n", monto, opcion)

	p.monto = monto
	p.opcion = opcion
	p.fechaPago = time.Now()
	p.id = time.Now().Format("060102150405") + "_" + suscripcion.id
	p.estado = "COMPLETADO"

	// Lógica de negocio: actualizar la suscripción
	suscripcion.Renovar() // <--- LLAMADA AL MÉTODO QUE DEFINIREMOS A CONTINUACIÓN

	fmt.Printf("✅ Pago ID %s registrado con éxito. Suscripción actualizada a Activa hasta %s.\n", p.id, suscripcion.fechaFin.Format("02-Jan-2006"))
	return nil
}

// VerificarPago simula la consulta del estado de un pago.
func (p *Pago) VerificarPago(pagoID string) bool {
	return p.estado == "COMPLETADO" && p.id == pagoID
}

// GetEstado devuelve el estado actual del pago.
func (p *Pago) GetEstado() string {
	return p.estado
}

// GetMonto devuelve el monto del pago.
func (p *Pago) GetMonto() float64 {
	return p.monto
}
