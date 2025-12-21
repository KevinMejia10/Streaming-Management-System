package content

import (
	"errors"
	"fmt"
)

var ErrContenidoNoEncontrado = errors.New("contenido no encontrado")

// Contenible actualizado con GetDescripcion
type Contenible interface {
	GetID() string
	GetTitulo() string
	GetDescripcion() string // <-- NUEVO: Ahora la descripción es accesible
	GetGenero() string
	GetDuracionTotal() float32
	Reproducir()
}

// --- ESTRUCTURA: Pelicula ---

type Pelicula struct {
	id          string
	titulo      string
	descripcion string
	genero      string
	director    string
	trailerLink string
	duracion    float32
}

func NuevaPelicula(id, titulo, descripcion, genero, director, trailerLink string, duracion float32) *Pelicula {
	return &Pelicula{
		id: id, titulo: titulo, descripcion: descripcion, genero: genero,
		director: director, trailerLink: trailerLink, duracion: duracion,
	}
}

func (p *Pelicula) GetID() string             { return p.id }
func (p *Pelicula) GetTitulo() string         { return p.titulo }
func (p *Pelicula) GetDescripcion() string    { return p.descripcion } // <-- Implementación
func (p *Pelicula) GetGenero() string         { return p.genero }
func (p *Pelicula) GetDuracionTotal() float32 { return p.duracion }
func (p *Pelicula) Reproducir() {
	fmt.Printf("▶️ Iniciando reproducción: Película '%s'\n", p.titulo)
}

// --- ESTRUCTURA: Serie ---

type Episodio struct {
	titulo   string
	duracion float32
}

type Serie struct {
	id          string
	titulo      string
	descripcion string
	genero      string
	temporadas  int
	episodios   []*Episodio
}

func NuevaSerie(id, titulo, descripcion, genero string, temporadas int) *Serie {
	return &Serie{
		id: id, titulo: titulo, descripcion: descripcion, genero: genero,
		temporadas: temporadas, episodios: make([]*Episodio, 0),
	}
}

func (s *Serie) GetID() string          { return s.id }
func (s *Serie) GetTitulo() string      { return s.titulo }
func (s *Serie) GetDescripcion() string { return s.descripcion } // <-- Implementación
func (s *Serie) GetGenero() string      { return s.genero }
func (s *Serie) GetDuracionTotal() float32 {
	var total float32
	for _, ep := range s.episodios {
		total += ep.duracion
	}
	return total
}
func (s *Serie) Reproducir() {
	fmt.Printf("▶️ Iniciando reproducción: Serie '%s' (Primer episodio)\n", s.titulo)
}

// Métodos adicionales de Serie
func (s *Serie) AgregarEpisodio(titulo string, duracion float32) {
	s.episodios = append(s.episodios, &Episodio{titulo: titulo, duracion: duracion})
}
func (s *Serie) ObtenerEpisodios() []*Episodio { return s.episodios }
