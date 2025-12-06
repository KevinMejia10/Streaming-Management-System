package auth

// Perfil representa un perfil de reproducción asociado a un Usuario.
type Perfil struct {
	id     int    // ID único del perfil
	nombre string // Nombre del perfil (e.g., "Adulto", "Niño")
}

// NuevoPerfil es un constructor para Perfil.
func NuevoPerfil(id int, nombre string) *Perfil {
	return &Perfil{
		id:     id,
		nombre: nombre,
	}
}

// GetID devuelve el ID del perfil.
func (p *Perfil) GetID() int {
	return p.id
}

// GetNombre devuelve el nombre del perfil.
func (p *Perfil) GetNombre() string {
	return p.nombre
}

// EditarNombre simula la modificación del nombre del perfil.
func (p *Perfil) EditarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// EliminarPerfil simula la eliminación del perfil (lógica a implementar).
func (p *Perfil) EliminarPerfil() {
	// Aquí iría la lógica de persistencia para eliminar.
}
