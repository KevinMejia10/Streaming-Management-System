package playback

import (
	"fmt" // Se requiere para el fmt.Printf en GuardarProgreso
	"streaming-system/pkg/content"
	"time"
)

// Visualizacion representa el registro de una única reproducción de contenido.
type Visualizacion struct {
	id              string
	contenido       content.Contenible
	fecha           time.Time
	guardarProgreso int // minuto de interrupción
}

// NuevoVisualizacion crea un nuevo registro de visualización.
func NuevoVisualizacion(id string, c content.Contenible) *Visualizacion {
	return &Visualizacion{
		id:              id,
		contenido:       c,
		fecha:           time.Now(),
		guardarProgreso: 0,
	}
}

// GetContenido devuelve el contenido que se está visualizando.
func (v *Visualizacion) GetContenido() content.Contenible {
	return v.contenido
}

// GuardarProgreso actualiza el minuto exacto donde se interrumpió la reproducción.
func (v *Visualizacion) GuardarProgreso(minuto int) {
	v.guardarProgreso = minuto
	fmt.Printf("💾 Progreso guardado: %s en el minuto %d.\n", v.contenido.GetTitulo(), minuto)
}

// ReproducirDesdePunto reanuda la reproducción.
func (v *Visualizacion) ReproducirDesdePunto() {
	if v.guardarProgreso > 0 {
		fmt.Printf("⏯️ Reanudando %s desde el minuto %d...\n", v.contenido.GetTitulo(), v.guardarProgreso)
	} else {
		v.contenido.Reproducir()
	}
}

// --- GETTERS AÑADIDOS PARA RESOLVER ERRORES DE ENCAPSULACIÓN ---

// GetGuardarProgreso devuelve el minuto de interrupción guardado (campo privado).
func (v *Visualizacion) GetGuardarProgreso() int {
	return v.guardarProgreso
}

// GetFecha devuelve la fecha y hora de la visualización (campo privado).
func (v *Visualizacion) GetFecha() time.Time {
	return v.fecha
}
