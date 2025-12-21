package content

// NOTA: ErrContenidoNoEncontrado y la interfaz Contenible
// se asumen definidos de forma única en pkg/content/content.go

// GestorDeContenido maneja el catálogo de películas y series.
type GestorDeContenido struct {
	catalogo map[string]Contenible
}

// NuevoGestorDeContenido crea e inicializa el catálogo.
func NuevoGestorDeContenido(contenidos []Contenible) *GestorDeContenido {
	g := &GestorDeContenido{
		catalogo: make(map[string]Contenible),
	}
	// Insertar contenidos cargados si existen
	for _, c := range contenidos {
		g.InsertarContenido(c)
	}
	return g
}

// InsertarContenido añade un elemento al catálogo.
func (g *GestorDeContenido) InsertarContenido(c Contenible) {
	g.catalogo[c.GetID()] = c
}

// ObtenerPorID recupera un contenido por su ID.
func (g *GestorDeContenido) ObtenerPorID(id string) (Contenible, error) {
	// Usamos la variable de error definida en content.go
	if c, ok := g.catalogo[id]; ok {
		return c, nil
	}
	return nil, ErrContenidoNoEncontrado
}

// ObtenerTodo devuelve todo el catálogo.
func (g *GestorDeContenido) ObtenerTodo() map[string]Contenible {
	return g.catalogo
}

// BorrarContenido elimina un contenido por su ID.
func (g *GestorDeContenido) BorrarContenido(id string) {
	delete(g.catalogo, id)
}

// ActualizarContenidoMetadata simula la actualización de metadata.
func (g *GestorDeContenido) ActualizarContenidoMetadata(id string, nuevoTitulo string) error {
	c, ok := g.catalogo[id]
	if !ok {
		return ErrContenidoNoEncontrado
	}

	switch v := c.(type) {
	case *Pelicula:
		v.titulo = nuevoTitulo
	case *Serie:
		v.titulo = nuevoTitulo
	default:
		// No se puede actualizar el tipo de contenido desconocido
	}

	return nil
}
