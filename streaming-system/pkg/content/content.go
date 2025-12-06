package content

import (
	"fmt"
	"time"
)

// Contenido representa el elemento base de una pieza de contenido.
type Contenido struct {
	id               string
	titulo           string
	descripcion      string
	duracion         float32 // en minutos
	genero           string
	fechaPublicacion time.Time
}

// Contenible define la interfaz para cualquier cosa que pueda ser reproducida.
type Contenible interface {
	GetID() string
	GetTitulo() string
	Reproducir()
}

// NuevoContenido crea una nueva instancia de Contenido.
func NuevoContenido(id, titulo, descripcion, genero string, duracion float32) Contenido {
	return Contenido{
		id:               id,
		titulo:           titulo,
		descripcion:      descripcion,
		duracion:         duracion,
		genero:           genero,
		fechaPublicacion: time.Now(),
	}
}

// GetID devuelve el ID del contenido.
func (c *Contenido) GetID() string {
	return c.id
}

// GetTitulo devuelve el título del contenido.
func (c *Contenido) GetTitulo() string {
	return c.titulo
}

// GetGenero devuelve el género del contenido.
func (c *Contenido) GetGenero() string {
	return c.genero
}

// GetFechaPublicacion devuelve la fecha de publicación.
func (c *Contenido) GetFechaPublicacion() time.Time {
	return c.fechaPublicacion
}

// GetDuracion devuelve la duración.
func (c *Contenido) GetDuracion() float32 {
	return c.duracion
}

// Pelicula extiende Contenido.
type Pelicula struct {
	Contenido // Composición/Herencia
	director  string
	trailer   string
}

// NuevaPelicula es el constructor de Pelicula.
func NuevaPelicula(id, titulo, desc, genero, director, trailer string, duracion float32) *Pelicula {
	return &Pelicula{
		Contenido: NuevoContenido(id, titulo, desc, genero, duracion),
		director:  director,
		trailer:   trailer,
	}
}

// Reproducir simula la reproducción de la película.
func (p *Pelicula) Reproducir() {
	fmt.Printf("🎥 Reproduciendo Película: %s (Dir: %s)\n", p.titulo, p.director)
}

// ReproducirTrailer simula la reproducción del trailer.
func (p *Pelicula) ReproducirTrailer() {
	fmt.Printf("▶️ Reproduciendo Trailer de: %s (Link: %s)\n", p.titulo, p.trailer)
}

// Episodio representa un episodio dentro de una Serie.
type Episodio struct {
	id       int
	titulo   string
	duracion float32
}

// NuevoEpisodio es el constructor de Episodio.
func NuevoEpisodio(id int, titulo string, duracion float32) *Episodio {
	return &Episodio{id: id, titulo: titulo, duracion: duracion}
}

// GetTitulo devuelve el título del episodio.
func (e *Episodio) GetTitulo() string {
	return e.titulo
}

// GetDuracion devuelve la duración del episodio.
func (e *Episodio) GetDuracion() float32 {
	return e.duracion
}

// Reproducir simula la reproducción del episodio.
func (e *Episodio) Reproducir() {
	fmt.Printf("▶️ Reproduciendo Episodio %d: %s\n", e.id, e.titulo)
}

// Serie extiende Contenido.
type Serie struct {
	Contenido  // Composición/Herencia
	temporadas int
	episodios  []*Episodio
}

// NuevaSerie es el constructor de Serie.
func NuevaSerie(id, titulo, desc, genero string, temporadas int) *Serie {
	return &Serie{
		Contenido:  NuevoContenido(id, titulo, desc, genero, 0), // Duración total no se usa aquí
		temporadas: temporadas,
		episodios:  make([]*Episodio, 0),
	}
}

// Reproducir simula la reproducción del primer episodio de la serie.
func (s *Serie) Reproducir() {
	fmt.Printf("📺 Iniciando reproducción de Serie: %s\n", s.titulo)
	if len(s.episodios) > 0 {
		s.episodios[0].Reproducir()
	} else {
		fmt.Println("⚠️ Sin episodios disponibles.")
	}
}

// AgregarEpisodio añade un nuevo episodio a la serie.
func (s *Serie) AgregarEpisodio(titulo string, duracion float32) *Episodio {
	nuevoEp := NuevoEpisodio(len(s.episodios)+1, titulo, duracion)
	s.episodios = append(s.episodios, nuevoEp)
	return nuevoEp
}

// ObtenerEpisodios devuelve la lista de episodios.
func (s *Serie) ObtenerEpisodios() []*Episodio {
	return s.episodios
}
