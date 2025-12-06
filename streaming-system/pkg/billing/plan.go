package billing

// Plan representa un plan de suscripción disponible.
type Plan struct {
	id          string
	nombre      string
	precio      float32
	mensualidad float32
}

// NuevoPlan es el constructor para Plan.
func NuevoPlan(id, nombre string, precio, mensualidad float32) *Plan {
	return &Plan{
		id:          id,
		nombre:      nombre,
		precio:      precio,
		mensualidad: mensualidad,
	}
}

// GetID devuelve el ID del plan.
func (p *Plan) GetID() string {
	return p.id
}

// GetNombre devuelve el nombre del plan.
func (p *Plan) GetNombre() string {
	return p.nombre
}

// GetPrecio devuelve el precio inicial.
func (p *Plan) GetPrecio() float32 {
	return p.precio
}

// GetMensualidad devuelve el precio recurrente.
func (p *Plan) GetMensualidad() float32 {
	return p.mensualidad
}

// ActualizarPrecio simula la actualización de precio.
func (p *Plan) ActualizarPrecio(nuevoPrecio float32) {
	p.precio = nuevoPrecio
}
