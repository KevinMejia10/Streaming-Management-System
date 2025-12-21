package billing

import (
	"time"
)

type Suscripcion struct {
	id          string
	plan        *Plan
	fechaInicio time.Time
	fechaFin    time.Time
	estado      string // ACTIVO, CANCELADO, EXPIRADO
}

func NuevaSuscripcion(id string, plan *Plan) *Suscripcion {
	return &Suscripcion{
		id:          id,
		plan:        plan,
		fechaInicio: time.Now(),
		fechaFin:    time.Now().AddDate(1, 0, 0),
		estado:      "ACTIVO",
	}
}

func RecreateSuscripcionFromDB(id string, plan *Plan, inicio, fin time.Time, estado string) *Suscripcion {
	return &Suscripcion{
		id:          id,
		plan:        plan,
		fechaInicio: inicio,
		fechaFin:    fin,
		estado:      estado,
	}
}

// --- NUEVOS MÉTODOS GETTERS (Añade estos para corregir los errores) ---

func (s *Suscripcion) GetID() string {
	return s.id
}

func (s *Suscripcion) GetFechaInicio() time.Time {
	return s.fechaInicio
}

func (s *Suscripcion) GetFechaFin() time.Time {
	return s.fechaFin
}

func (s *Suscripcion) GetEstado() string {
	return s.estado
}

func (s *Suscripcion) GetPlan() *Plan {
	return s.plan
}

// --- MÉTODOS DE LÓGICA ---

func (s *Suscripcion) EstaActiva() bool {
	return s.estado == "ACTIVO" && time.Now().Before(s.fechaFin)
}

func (s *Suscripcion) Renovar() {
	if s.EstaActiva() {
		s.fechaFin = s.fechaFin.AddDate(1, 0, 0)
	} else {
		s.fechaInicio = time.Now()
		s.fechaFin = s.fechaInicio.AddDate(1, 0, 0)
	}
	s.estado = "ACTIVO"
}
