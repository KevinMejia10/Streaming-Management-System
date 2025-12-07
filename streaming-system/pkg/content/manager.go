package content

import (
	"errors"
	"time"
)

// GestorDeContenido que maneja todo el inventario de contenido.
type GestorDeContenido struct {
	catalogo map[string]Contenible
}

// NuevoGestorDeContenido inicializa el gestor.
func NuevoGestorDeContenido() *GestorDeContenido {
	return &GestorDeContenido{
		catalogo: make(map[string]Contenible),
	}
}

// ObtenerTodo el catálogo.
func (g *GestorDeContenido) ObtenerTodo() map[string]Contenible {
	return g.catalogo
}

// InsertarContenido añade un nuevo Contenible al catálogo.
func (g *GestorDeContenido) InsertarContenido(c Contenible) {
	g.catalogo[c.GetID()] = c
}

// BorrarContenido elimina un Contenible del catálogo.
func (g *GestorDeContenido) BorrarContenido(contenidoID string) {
	delete(g.catalogo, contenidoID)
}

// ActualizarContenidoMetadata simula la actualización de un título.
func (g *GestorDeContenido) ActualizarContenidoMetadata(contenidoID string, nuevoTitulo string) error {
	c, ok := g.catalogo[contenidoID]
	if !ok {
		return errors.New("contenido no encontrado")
	}

	// Simulación de actualización: solo modificamos el campo 'titulo' si es de tipo Contenido
	switch v := c.(type) {
	case *Pelicula:
		v.titulo = nuevoTitulo
	case *Serie:
		v.titulo = nuevoTitulo
	default:
		return errors.New("tipo de contenido desconocido para actualizar")
	}
	return nil
}

// BuscarContenido busca por ID o título.
func (g *GestorDeContenido) BuscarContenido(query string) (Contenible, error) {
	// Búsqueda por ID
	if c, ok := g.catalogo[query]; ok {
		return c, nil
	}

	// Búsqueda simple por título (primera coincidencia)
	for _, c := range g.catalogo {
		if c.GetTitulo() == query {
			return c, nil
		}
	}

	return nil, errors.New("contenido no encontrado")
}

// FiltrarContenido permite buscar por género y fecha de publicación.
func (g *GestorDeContenido) FiltrarContenido(genero string, fecha time.Time) []Contenible {
	resultados := make([]Contenible, 0)
	for _, c := range g.catalogo {
		// En Go, necesitamos un "type assertion" o una función para acceder a los campos de Contenido
		// En este diseño, usaremos las funciones accesoras de la interfaz Contenible si se definieran.
		// Como solo Contenido tiene GetGenero/GetFecha, asumimos que todos los Contenibles son Pelicula/Serie.

		var cumpleGenero bool
		var cumpleFecha bool

		switch v := c.(type) {
		case *Pelicula:
			cumpleGenero = genero == "" || v.Contenido.GetGenero() == genero
			cumpleFecha = fecha.IsZero() || v.Contenido.GetFechaPublicacion().After(fecha)
		case *Serie:
			cumpleGenero = genero == "" || v.Contenido.GetGenero() == genero
			cumpleFecha = fecha.IsZero() || v.Contenido.GetFechaPublicacion().After(fecha)
		}

		if cumpleGenero && cumpleFecha {
			resultados = append(resultados, c)
		}
	}
	return resultados
}
