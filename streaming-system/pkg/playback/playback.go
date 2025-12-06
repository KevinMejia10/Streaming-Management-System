package playback

import "fmt"

// HistorialReproduccion almacena las Visualizaciones para un Perfil.
type HistorialReproduccion struct {
	visualizaciones map[string]*Visualizacion // Clave: ContenidoID
}

// NuevoHistorialReproduccion crea un nuevo historial.
func NuevoHistorialReproduccion() *HistorialReproduccion {
	return &HistorialReproduccion{
		visualizaciones: make(map[string]*Visualizacion),
	}
}

// AgregarVisualizacion registra o actualiza una visualización en el historial.
func (h *HistorialReproduccion) AgregarVisualizacion(v *Visualizacion) {
	h.visualizaciones[v.contenido.GetID()] = v
}

// GetVisualizaciones devuelve la lista de visualizaciones.
func (h *HistorialReproduccion) GetVisualizaciones() map[string]*Visualizacion {
	return h.visualizaciones
}

// ObtenerUltimaVisualizacionPorContenido devuelve la visualización de un contenido.
func (h *HistorialReproduccion) ObtenerUltimaVisualizacionPorContenido(contenidoID string) *Visualizacion {
	return h.visualizaciones[contenidoID]
}

// FiltrarPorFecha simula la filtración del historial por fecha.
func (h *HistorialReproduccion) FiltrarPorFecha() {
	// [Lógica pendiente de implementar en el diagrama]
	fmt.Println("Filtrando historial por fecha (Funcionalidad no implementada en el diagrama)")
}

// EliminarHistorialFinalizado elimina entradas con progreso completo.
func (h *HistorialReproduccion) EliminarHistorialFinalizado() {
	// [Lógica pendiente de implementar en el diagrama]
	fmt.Println("Eliminando historial de visualizaciones finalizadas (Funcionalidad no implementada en el diagrama)")
}
