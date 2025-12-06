package billing

import "time"

// Suscripcion representa la asociación de un Usuario con un Plan.
type Suscripcion struct {
	id          string
	plan        *Plan
	fechaInicio time.Time
	fechaFin    time.Time
	estado      string // "Activa", "Inactiva", "Pendiente"
}

// NuevaSuscripcion crea una nueva suscripción.
func NuevaSuscripcion(id string, plan *Plan) *Suscripcion {
	return &Suscripcion{
		id:          id,
		plan:        plan,
		fechaInicio: time.Now(),
		fechaFin:    time.Now().AddDate(0, 1, 0), // Un mes por defecto
		estado:      "Pendiente",
	}
}

// GetPlan devuelve el plan asociado.
func (s *Suscripcion) GetPlan() *Plan {
	return s.plan
}

// EstaActiva verifica si la suscripción está activa y no ha expirado.
func (s *Suscripcion) EstaActiva() bool {
	return s.estado == "Activa" && time.Now().Before(s.fechaFin)
}

// Renovar simula la renovación de la suscripción por un periodo.
func (s *Suscripcion) Renovar() {
	s.fechaFin = s.fechaFin.AddDate(0, 1, 0)
	s.estado = "Activa"
}

// Cancelar simula la cancelación de la suscripción.
func (s *Suscripcion) Cancelar() {
	s.estado = "Inactiva"
	// Podríamos no modificar fechaFin para permitir acceso hasta el final del ciclo
}
