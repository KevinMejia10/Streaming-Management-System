package billing

// Plan representa un plan de suscripción disponible.
type Plan struct {
	id     string
	nombre string
	precio float64
	costo  float64
}

// NuevoPlan es el constructor para Plan.
func NuevoPlan(id, nombre string, precio, costo float64) *Plan {
	return &Plan{
		id:     id,
		nombre: nombre,
		precio: precio,
		costo:  costo,
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

// GetPrecio devuelve el precio del plan.
func (p *Plan) GetPrecio() float64 {
	return p.precio
}
