package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// --- MÓDULO DE ERRORES PERSONALIZADOS ---

// Definición de errores comunes para el sistema.
var (
	ErrCredencialesInvalidas = errors.New("autenticacion: credenciales invalidas")
	ErrUsuarioNoEncontrado   = errors.New("usuario: no encontrado")
	ErrPlanInvalido          = errors.New("suscripcion: plan invalido")
	ErrContenidoNoEncontrado = errors.New("gestor_contenido: contenido no encontrado")
	ErrNoSuscripcionActiva   = errors.New("autorizacion: el usuario no tiene una suscripcion activa")
	ErrPagoFallido           = errors.New("pago: la simulacion de pago ha fallado")
	ErrPerfilNoEncontrado    = errors.New("perfil: no encontrado")
)

// --- INTERFACES GLOBALES ---

// OpcionDePago define el contrato para cualquier mecanismo de pago (Tarjeta, PayPal, etc.).
type OpcionDePago interface {
	ProcesarPago(monto float64) error
	GetNombre() string
}

// Pagable define el comportamiento para las entidades que pueden ser pagadas (e.g., Suscripcion).
type Pagable interface {
	GetMonto() float64
	RealizarPago(opcion OpcionDePago) error
}

// Reproducible define el comportamiento base para el Contenido que puede ser reproducido.
type Reproducible interface {
	GetID() int
	GetTitulo() string
	Reproducir(p *Perfil)
}

// --- ESTRUCTURAS DE DATOS (CLASES) ---

// ------------------------------------
// 🎯 MÓDULO DE AUTENTICACIÓN Y USUARIO
// ------------------------------------

// Autenticacion representa el token de un usuario activo.
type Autenticacion struct {
	token      string
	expiraDate time.Time
	usuarioID  int
}

// generarToken simula la creación de un nuevo token de sesión.
func (a *Autenticacion) generarToken(u *Usuario) {
	a.token = fmt.Sprintf("token-%d-%d", u.id, time.Now().Unix())
	a.expiraDate = time.Now().Add(12 * time.Hour) // Token válido por 12 horas
	a.usuarioID = u.id
}

// revocarToken simula la anulación del token actual.
func (a *Autenticacion) revocarToken() {
	a.token = ""
	a.expiraDate = time.Time{}
	a.usuarioID = 0
}

// HistorialReproduccion almacena las visualizaciones de un perfil.
type HistorialReproduccion struct {
	id              int
	visualizaciones map[int]*Visualizacion // Clave: ID de Visualizacion
}

// agregarVisualizacion agrega un registro al historial, asegurando unicidad por contenido.
func (h *HistorialReproduccion) agregarVisualizacion(v *Visualizacion) {
	// Usamos el ContentID como clave para reemplazar si ya existe (actualización de progreso).
	h.visualizaciones[v.contenidoID] = v
}

// Visualizacion representa el progreso de reproducción de un contenido por un perfil.
type Visualizacion struct {
	id          int
	fecha       time.Time
	minuto      float64 // Minuto de interrupción.
	contenidoID int
	perfilID    int
	usuarioID   int
}

// guardarProgreso actualiza el progreso de visualización.
func (v *Visualizacion) guardarProgreso(minuto float64) {
	v.minuto = minuto
	v.fecha = time.Now()
	fmt.Printf(" [💾] Progreso de '%s' guardado en %.1f min.\n",
		Sistema.Gestor.GetContenidoPorID(v.contenidoID).GetTitulo(), minuto)
}

// reproducirDesdePunto simula la reanudación de la reproducción.
func (v *Visualizacion) reproducirDesdePunto() {
	fmt.Printf(" [▶️] Reanudando '%s' desde el minuto %.1f...\n",
		Sistema.Gestor.GetContenidoPorID(v.contenidoID).GetTitulo(), v.minuto)
	// Simulación de reproducción completa
	v.guardarProgreso(9999.0) // Valor alto para simular "visto"
}

// Perfil representa un perfil dentro de la cuenta de un usuario.
type Perfil struct {
	id     int
	nombre string
	avatar string
}

// crearPerfil inicializa un nuevo perfil.
func (p *Perfil) crearPerfil(id int, nombre, avatar string) {
	p.id = id
	p.nombre = nombre
	p.avatar = avatar
}

// editarPerfil actualiza el nombre y avatar.
func (p *Perfil) editarPerfil(nombre, avatar string) {
	p.nombre = nombre
	p.avatar = avatar
	fmt.Printf(" [👤] Perfil '%s' actualizado.\n", nombre)
}

// eliminarPerfil simula la eliminación (la lógica real estaría en Usuario).
func (*Perfil) eliminarPerfil() {
	// Lógica de eliminación...
}

// Usuario representa una cuenta de usuario del sistema de streaming.
type Usuario struct {
	id                    int
	nombre                string
	email                 string
	contraseniaHash       string // Simulación de contraseña hasheada
	perfiles              map[int]*Perfil
	autenticacion         *Autenticacion
	suscripcion           *Suscripcion
	historialReproduccion *HistorialReproduccion
}

// nuevoUsuario crea una nueva instancia de Usuario.
func nuevoUsuario(id int, nombre, email, contrasenia string) *Usuario {
	return &Usuario{
		id:                    id,
		nombre:                nombre,
		email:                 email,
		contraseniaHash:       hashContrasenia(contrasenia), // Simulación de hashing
		perfiles:              make(map[int]*Perfil),
		autenticacion:         &Autenticacion{},
		historialReproduccion: &HistorialReproduccion{id: id, visualizaciones: make(map[int]*Visualizacion)},
	}
}

// hashContrasenia simula el hashing de una contraseña.
func hashContrasenia(contrasenia string) string {
	return "hash_" + contrasenia // Lógica real usaría bcrypt/sha256
}

// verificarContrasenia simula la verificación de la contraseña hasheada.
func (u *Usuario) verificarContrasenia(contrasenia string) bool {
	return u.contraseniaHash == hashContrasenia(contrasenia)
}

// iniciarSesion verifica credenciales y genera un token.
func (u *Usuario) iniciarSesion(email, contrasenia string) (*Autenticacion, error) {
	if u.email != email || !u.verificarContrasenia(contrasenia) {
		return nil, ErrCredencialesInvalidas
	}
	u.autenticacion.generarToken(u)
	fmt.Printf(" [🔑] Inicio de sesión exitoso para %s.\n", u.nombre)
	return u.autenticacion, nil
}

// cerrarSesion revoca el token.
func (u *Usuario) cerrarSesion() {
	u.autenticacion.revocarToken()
	fmt.Printf(" [🚪] Sesión cerrada para %s.\n", u.nombre)
}

// crearPerfil añade un nuevo perfil al usuario.
func (u *Usuario) crearPerfil(nombre, avatar string) *Perfil {
	nextID := len(u.perfiles) + 1
	p := &Perfil{}
	p.crearPerfil(nextID, nombre, avatar)
	u.perfiles[nextID] = p
	fmt.Printf(" [👤] Perfil '%s' creado exitosamente.\n", p.nombre)
	return p
}

// verPerfiles lista los perfiles del usuario.
func (u *Usuario) verPerfiles() map[int]*Perfil {
	return u.perfiles
}

// verHistorialReproduccion devuelve el historial del usuario.
func (u *Usuario) verHistorialReproduccion() *HistorialReproduccion {
	return u.historialReproduccion
}

// verificarSuscripcionActiva comprueba si la suscripción está vigente.
func (u *Usuario) verificarSuscripcionActiva() bool {
	if u.suscripcion == nil {
		return false
	}
	return u.suscripcion.estaActiva()
}

// getPerfilPorID obtiene un perfil por su ID.
func (u *Usuario) getPerfilPorID(id int) (*Perfil, error) {
	p, ok := u.perfiles[id]
	if !ok {
		return nil, ErrPerfilNoEncontrado
	}
	return p, nil
}

// -----------------------------------
// 💰 MÓDULO DE SUSCRIPCIONES Y PAGOS
// -----------------------------------

// Plan define las características y el costo de un plan de suscripción.
type Plan struct {
	id          int
	nombre      string
	precio      float64
	maxDisposit int
	calidad     string
}

// actualizarDetalles cambia los atributos del plan.
func (p *Plan) actualizarDetalles(nombre string, precio float64, maxDisposit int, calidad string) {
	p.nombre = nombre
	p.precio = precio
	p.maxDisposit = maxDisposit
	p.calidad = calidad
}

// Suscripcion representa la suscripción activa de un usuario a un Plan.
type Suscripcion struct {
	id          int
	fechaInicio time.Time
	fechaFin    time.Time
	estaActiva  bool
	plan        *Plan
	usuarioID   int
}

// renovar extiende la duración de la suscripción.
func (s *Suscripcion) renovar(duracion time.Duration) {
	s.fechaFin = s.fechaFin.Add(duracion)
	s.estaActiva = true
}

// cancelar establece la suscripción como inactiva (a menudo al final del ciclo).
func (s *Suscripcion) cancelar() {
	s.estaActiva = false
	fmt.Printf(" [🛑] Suscripción cancelada. Válida hasta %s.\n", s.fechaFin.Format("02-Jan-2006"))
}

// estaActiva comprueba si la suscripción está vigente en la fecha actual.
func (s *Suscripcion) estaActiva() bool {
	// Verificar que el campo `estaActiva` sea verdadero Y que la fecha de fin sea posterior a hoy.
	return s.estaActiva && time.Now().Before(s.fechaFin)
}

// Pago representa un registro de pago asociado a una suscripción.
type Pago struct {
	id             int
	fecha          time.Time
	monto          float64
	metodo         string // E.g., "Tarjeta", "PayPal"
	procesado      bool   // Indica si la transacción fue exitosa
	suscripcionRef *Suscripcion
}

// GetMonto implementa Pagable.
func (p *Pago) GetMonto() float64 {
	return p.monto
}

// RealizarPago simula el proceso de pago e implementa Pagable.
func (p *Pago) RealizarPago(opcion OpcionDePago) error {
	fmt.Printf(" [💳] Procesando pago de %.2f USD con %s...\n", p.monto, opcion.GetNombre())
	if err := opcion.ProcesarPago(p.monto); err != nil {
		p.procesado = false
		return err // Reenviar error
	}
	p.procesado = true
	p.fecha = time.Now()
	p.metodo = opcion.GetNombre()
	// Actualizar Suscripcion en caso de éxito
	p.suscripcionRef.renovar(30 * 24 * time.Hour) // Renovar por 30 días
	fmt.Printf(" [✅] Pago exitoso. Suscripción renovada hasta %s.\n", p.suscripcionRef.fechaFin.Format("02-Jan-2006"))
	return nil
}

// registrarPago crea y ejecuta el pago.
func (p *Pago) registrarPago(monto float64, suscripcion *Suscripcion) {
	p.monto = monto
	p.suscripcionRef = suscripcion
}

// ------------------------------------
// 🎬 MÓDULO DE CONTENIDO Y REPRODUCCIÓN
// ------------------------------------

// Contenido es la estructura base para todo el contenido multimedia.
type Contenido struct {
	id               int
	titulo           string
	descripcion      string
	genero           string
	duracion         float64 // En minutos
	fechaPublicacion time.Time
}

// Reproducir simula la reproducción del contenido.
func (c *Contenido) Reproducir(p *Perfil) {
	if !Sistema.GetUsuarioActual().verificarSuscripcionActiva() {
		fmt.Printf(" [❌] ERROR: No puedes reproducir '%s'. %s\n", c.titulo, ErrNoSuscripcionActiva)
		return
	}

	historial := Sistema.GetUsuarioActual().verHistorialReproduccion()
	var visualizacion *Visualizacion
	var ok bool

	// 1. Buscar si hay progreso previo
	for _, v := range historial.visualizaciones {
		if v.contenidoID == c.id && v.perfilID == p.id {
			visualizacion = v
			ok = true
			break
		}
	}

	if !ok {
		// 2. Crear nueva visualización si no existe
		fmt.Printf(" [🆕] Iniciando reproducción de '%s' para el perfil %s.\n", c.titulo, p.nombre)
		visualizacion = &Visualizacion{
			id:          len(historial.visualizaciones) + 1,
			fecha:       time.Now(),
			minuto:      0.0,
			contenidoID: c.id,
			perfilID:    p.id,
			usuarioID:   p.id, // En realidad, sería el ID del usuario
		}
		historial.agregarVisualizacion(visualizacion)

	} else if visualizacion.minuto > 0 && visualizacion.minuto < c.duracion {
		// 3. Reanudar desde el punto si no está completado
		visualizacion.reproducirDesdePunto()
		return
	}

	// 4. Simulación de reproducción y guardado de progreso
	fmt.Printf(" [▶️] Reproduciendo '%s' (%s - %.1f min)...\n", c.titulo, c.genero, c.duracion)
	simuladoProgreso := 5.5
	if c.duracion-simuladoProgreso < 1.0 {
		simuladoProgreso = c.duracion // Simular finalización
	}
	visualizacion.guardarProgreso(simuladoProgreso)

	// Simular si el usuario terminó de verlo para eliminar la "simulación" de progreso
	if visualizacion.minuto >= c.duracion {
		visualizacion.minuto = c.duracion
		fmt.Println(" [👍] Contenido marcado como visto.")
	}

}

// GetID implementa Reproducible.
func (c *Contenido) GetID() int { return c.id }

// GetTitulo implementa Reproducible.
func (c *Contenido) GetTitulo() string { return c.titulo }

// Pelicula es un tipo de contenido que extiende la estructura base.
type Pelicula struct {
	Contenido
	director string
	trailer  string
}

// reproducirTrailer simula la reproducción de un trailer.
func (p *Pelicula) reproducirTrailer() {
	fmt.Printf(" [🎞️] Reproduciendo trailer de: %s (Director: %s)\n", p.titulo, p.director)
}

// Serie es un tipo de contenido contenedor de Episodios.
type Serie struct {
	Contenido
	temporadas int
	episodios  map[int][]*Episodio // Clave: Número de temporada
}

// Episodio es la unidad de reproducción para una Serie.
type Episodio struct {
	titulo       string
	duracion     float64
	numEpisodio  int
	numTemporada int
}

// Reproducir simula la reproducción de un episodio.
func (e *Episodio) Reproducir(p *Perfil) {
	fmt.Printf(" [▶️] Reproduciendo S%d E%d: %s (%.1f min)\n", e.numTemporada, e.numEpisodio, e.titulo, e.duracion)
	// La lógica de historial debería actualizarse a través de la Serie/Visualizacion
}

// agregarEpisodio añade un episodio a una temporada.
func (s *Serie) agregarEpisodio(e *Episodio) {
	if s.episodios == nil {
		s.episodios = make(map[int][]*Episodio)
	}
	if _, ok := s.episodios[e.numTemporada]; !ok {
		s.episodios[e.numTemporada] = make([]*Episodio, 0)
		s.temporadas++
	}
	s.episodios[e.numTemporada] = append(s.episodios[e.numTemporada], e)

	// Actualizar la duración total de la serie
	s.duracion += e.duracion
}

// listarEpisodios devuelve los episodios por temporada.
func (s *Serie) listarEpisodios(temporada int) []*Episodio {
	return s.episodios[temporada]
}

// ------------------------------------
// ⚙️ GESTOR DE CONTENIDO
// ------------------------------------

// GestorDeContenido maneja todo el catálogo de contenido del sistema.
type GestorDeContenido struct {
	catalogo map[int]Reproducible // Clave: ID de contenido
}

// GetContenidoPorID devuelve contenido por su ID.
func (g *GestorDeContenido) GetContenidoPorID(id int) Reproducible {
	return g.catalogo[id]
}

// insertarContenido agrega un nuevo título al catálogo.
func (g *GestorDeContenido) insertarContenido(c Reproducible) {
	g.catalogo[c.GetID()] = c
	fmt.Printf(" [➕] Contenido '%s' agregado al catálogo.\n", c.GetTitulo())
}

// actualizarContenido simula la actualización de metadata.
func (*GestorDeContenido) actualizarContenido(id int, titulo, descripcion, genero string, duracion float64, fecha time.Time, esPelicula bool) {
	// Lógica para encontrar y actualizar el contenido...
}

// buscarContenido busca contenido por título (sensible a mayúsculas/minúsculas).
func (g *GestorDeContenido) buscarContenido(query string) []Reproducible {
	var resultados []Reproducible
	lowerQuery := strings.ToLower(query)
	for _, c := range g.catalogo {
		// Búsqueda en Película
		if p, ok := c.(*Pelicula); ok && strings.Contains(strings.ToLower(p.titulo), lowerQuery) {
			resultados = append(resultados, p)
			continue
		}
		// Búsqueda en Serie
		if s, ok := c.(*Serie); ok && strings.Contains(strings.ToLower(s.titulo), lowerQuery) {
			resultados = append(resultados, s)
		}
	}
	return resultados
}

// filtrarContenido filtra por género y/o fecha de publicación.
func (g *GestorDeContenido) filtrarContenido(genero string, anio int) []Reproducible {
	var resultados []Reproducible
	for _, c := range g.catalogo {
		match := true
		if genero != "" && c.(*Contenido).genero != genero {
			match = false
		}
		if anio != 0 && c.(*Contenido).fechaPublicacion.Year() != anio {
			match = false
		}
		if match {
			resultados = append(resultados, c)
		}
	}
	return resultados
}

// listarTodo lista todo el catálogo.
func (g *GestorDeContenido) listarTodo() []Reproducible {
	var lista []Reproducible
	for _, c := range g.catalogo {
		lista = append(lista, c)
	}
	return lista
}

// --- IMPLEMENTACIONES DE OPCIONDEPAGO ---

// PagoTarjeta simula una opción de pago con tarjeta.
type PagoTarjeta struct{}

// ProcesarPago simula la validación de tarjeta.
func (*PagoTarjeta) ProcesarPago(monto float64) error {
	// Simulación de fallo
	if monto > 100 {
		return ErrPagoFallido
	}
	// Lógica de gateway de pago
	return nil
}

// GetNombre implementa OpcionDePago.
func (*PagoTarjeta) GetNombre() string { return "Tarjeta de Crédito/Débito" }

// PagoPayPal simula una opción de pago con PayPal.
type PagoPayPal struct{}

// ProcesarPago simula la validación de PayPal.
func (*PagoPayPal) ProcesarPago(monto float64) error {
	// Simulación: siempre exitoso para el demo
	return nil
}

// GetNombre implementa OpcionDePago.
func (*PagoPayPal) GetNombre() string { return "PayPal" }

// --- ESTRUCTURA DEL SISTEMA Y DATOS MOCK ---

// SistemaStreaming contiene el estado global de la aplicación (usuarios, gestores).
type SistemaStreaming struct {
	Usuarios map[int]*Usuario
	Planes   map[int]*Plan
	Gestor   *GestorDeContenido
	// Estado de la sesión actual
	usuarioActual *Usuario
}

// NuevoSistema inicializa el sistema con datos de ejemplo.
func NuevoSistema() *SistemaStreaming {
	s := &SistemaStreaming{
		Usuarios: make(map[int]*Usuario),
		Planes:   make(map[int]*Plan),
		Gestor:   &GestorDeContenido{catalogo: make(map[int]Reproducible)},
	}
	s.inicializarPlanes()
	s.inicializarContenido()
	return s
}

// inicializarPlanes crea planes de suscripción mock.
func (s *SistemaStreaming) inicializarPlanes() {
	s.Planes[1] = &Plan{id: 1, nombre: "Básico", precio: 7.99, maxDisposit: 1, calidad: "SD"}
	s.Planes[2] = &Plan{id: 2, nombre: "Estándar", precio: 12.99, maxDisposit: 2, calidad: "HD"}
	s.Planes[3] = &Plan{id: 3, nombre: "Premium", precio: 17.99, maxDisposit: 4, calidad: "4K"}
}

// inicializarContenido llena el catálogo mock.
func (s *SistemaStreaming) inicializarContenido() {
	// Películas
	p1 := &Pelicula{
		Contenido: Contenido{id: 101, titulo: "La Odisea Estelar", descripcion: "Viaje épico a la galaxia.", genero: "Ciencia Ficción", duracion: 125.0, fechaPublicacion: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		director:  "A. Clarke", trailer: "link/trailer/101",
	}
	s.Gestor.insertarContenido(p1)

	p2 := &Pelicula{
		Contenido: Contenido{id: 102, titulo: "El Misterio del Faro", descripcion: "Thriller psicológico.", genero: "Thriller", duracion: 98.5, fechaPublicacion: time.Date(2022, 11, 20, 0, 0, 0, 0, time.UTC)},
		director:  "B. Hitchcock", trailer: "link/trailer/102",
	}
	s.Gestor.insertarContenido(p2)

	// Series
	ser1 := &Serie{
		Contenido: Contenido{id: 201, titulo: "Los Herederos del Dragón", descripcion: "Fantasía medieval y guerra.", genero: "Fantasía", duracion: 0.0, fechaPublicacion: time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)},
	}

	e1s1 := &Episodio{titulo: "El Llamado", duracion: 55.0, numEpisodio: 1, numTemporada: 1}
	e2s1 := &Episodio{titulo: "El Juramento", duracion: 62.0, numEpisodio: 2, numTemporada: 1}
	ser1.agregarEpisodio(e1s1)
	ser1.agregarEpisodio(e2s1)

	ser1.agregarEpisodio(&Episodio{titulo: "La Traición", duracion: 70.0, numEpisodio: 1, numTemporada: 2})

	s.Gestor.insertarContenido(ser1)
}

// GetUsuarioPorEmail busca un usuario.
func (s *SistemaStreaming) GetUsuarioPorEmail(email string) (*Usuario, error) {
	for _, u := range s.Usuarios {
		if u.email == email {
			return u, nil
		}
	}
	return nil, ErrUsuarioNoEncontrado
}

// SetUsuarioActual establece el usuario que ha iniciado sesión.
func (s *SistemaStreaming) SetUsuarioActual(u *Usuario) {
	s.usuarioActual = u
}

// GetUsuarioActual devuelve el usuario en sesión.
func (s *SistemaStreaming) GetUsuarioActual() *Usuario {
	return s.usuarioActual
}

// RegistrarUsuario agrega un nuevo usuario al sistema.
func (s *SistemaStreaming) RegistrarUsuario(nombre, email, contrasenia string) (*Usuario, error) {
	if _, err := s.GetUsuarioPorEmail(email); err == nil {
		return nil, errors.New("registro: el email ya está en uso")
	}

	nextID := len(s.Usuarios) + 1
	u := nuevoUsuario(nextID, nombre, email, contrasenia)
	s.Usuarios[nextID] = u
	return u, nil
}

// --- FUNCIÓN PRINCIPAL Y MENÚ DE CONSOLA ---

// Variable global para el sistema de streaming.
var Sistema *SistemaStreaming

func main() {
	Sistema = NuevoSistema()
	fmt.Println("===================================================")
	fmt.Println("📺 Sistema de Gestión de Streaming (POO & Go Idiomatic)")
	fmt.Println("===================================================")

	menuPrincipal()
}

func menuPrincipal() {
	var opcion int
	for {
		if Sistema.GetUsuarioActual() == nil {
			fmt.Println("\n--- MENÚ PRINCIPAL ---")
			fmt.Println("1. Registrar Nuevo Usuario")
			fmt.Println("2. Iniciar Sesión")
			fmt.Println("3. Salir")
			fmt.Print("Elige una opción: ")
			fmt.Scan(&opcion)

			switch opcion {
			case 1:
				ejecutarRegistro()
			case 2:
				ejecutarInicioSesion()
			case 3:
				fmt.Println("\nAdiós. ¡Vuelve pronto!")
				return
			default:
				fmt.Println("Opción no válida.")
			}
		} else {
			menuUsuario(Sistema.GetUsuarioActual())
		}
	}
}

func ejecutarRegistro() {
	var nombre, email, contrasenia string
	fmt.Print("Nombre: ")
	fmt.Scan(&nombre)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Contraseña: ")
	fmt.Scan(&contrasenia)

	u, err := Sistema.RegistrarUsuario(nombre, email, contrasenia)
	if err != nil {
		fmt.Printf(" [❌] Error de registro: %v\n", err)
	} else {
		fmt.Printf(" [👍] Usuario %s registrado con éxito. ID: %d\n", u.nombre, u.id)
	}
}

func ejecutarInicioSesion() {
	var email, contrasenia string
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Contraseña: ")
	fmt.Scan(&contrasenia)

	u, err := Sistema.GetUsuarioPorEmail(email)
	if err != nil || u == nil {
		fmt.Printf(" [❌] Error de Inicio de Sesión: %v\n", ErrCredencialesInvalidas)
		return
	}

	_, err = u.iniciarSesion(email, contrasenia)
	if err != nil {
		fmt.Printf(" [❌] Error de Inicio de Sesión: %v\n", err)
		return
	}
	Sistema.SetUsuarioActual(u)
}

func menuUsuario(u *Usuario) {
	var opcion int
	for Sistema.GetUsuarioActual() != nil {
		fmt.Printf("\n--- MENÚ DE USUARIO (%s) ---\n", u.nombre)
		fmt.Printf("Estado de Suscripción: **%s**\n", func() string {
			if u.verificarSuscripcionActiva() {
				return fmt.Sprintf("ACTIVA (Plan: %s)", u.suscripcion.plan.nombre)
			}
			return "INACTIVA"
		}())
		fmt.Println("1. Gestionar Suscripción y Pago")
		fmt.Println("2. Crear Perfil")
		fmt.Println("3. Seleccionar Perfil y Ver Contenido/Historial")
		fmt.Println("4. Cerrar Sesión")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			menuSuscripcion(u)
		case 2:
			ejecutarCrearPerfil(u)
		case 3:
			menuReproduccion(u)
		case 4:
			u.cerrarSesion()
			Sistema.SetUsuarioActual(nil)
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Función auxiliar para obtener entrada de texto completa (que podría incluir espacios).
func obtenerEntradaTexto(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln() // Limpiar buffer después de Scan
	fmt.Scanln(&input)
	return input
}

func ejecutarCrearPerfil(u *Usuario) {
	if len(u.verPerfiles()) >= 5 {
		fmt.Println(" [⚠️] Has alcanzado el límite de 5 perfiles.")
		return
	}
	var nombre, avatar string
	fmt.Print("Nombre del nuevo perfil: ")
	fmt.Scan(&nombre)
	fmt.Print("Avatar (ej: 'Dragon', 'Robot'): ")
	fmt.Scan(&avatar)
	u.crearPerfil(nombre, avatar)
}

func menuSuscripcion(u *Usuario) {
	if u.suscripcion != nil && u.verificarSuscripcionActiva() {
		fmt.Println(" [ℹ️] Ya tienes una suscripción activa.")
		return
	}

	fmt.Println("\n--- PLANES DE SUSCRIPCIÓN ---")
	for id, p := range Sistema.Planes {
		fmt.Printf("%d. %s (%.2f USD/mes) - Calidad: %s\n", id, p.nombre, p.precio, p.calidad)
	}

	var planID int
	fmt.Print("Elige el ID del plan a suscribir/renovar: ")
	fmt.Scan(&planID)

	planElegido, ok := Sistema.Planes[planID]
	if !ok {
		fmt.Printf(" [❌] Error: %v\n", ErrPlanInvalido)
		return
	}

	if u.suscripcion == nil {
		// Crear nueva suscripción (simulando que el ID es incremental)
		u.suscripcion = &Suscripcion{
			id:        len(Sistema.Usuarios) + 1000,
			plan:      planElegido,
			usuarioID: u.id,
			// Los campos de fechaInicio/Fin se llenarán tras el pago
		}
	} else {
		// Asignar el nuevo plan para renovación/cambio de plan
		u.suscripcion.plan = planElegido
	}

	// Crear el registro de pago
	pago := &Pago{}
	pago.registrarPago(planElegido.precio, u.suscripcion)

	fmt.Println("\n--- OPCIONES DE PAGO ---")
	fmt.Println("1. Tarjeta de Crédito/Débito")
	fmt.Println("2. PayPal")
	var pagoOpcion int
	fmt.Print("Selecciona la opción de pago: ")
	fmt.Scan(&pagoOpcion)

	var opcionDePago OpcionDePago
	switch pagoOpcion {
	case 1:
		opcionDePago = &PagoTarjeta{}
	case 2:
		opcionDePago = &PagoPayPal{}
	default:
		fmt.Println(" [❌] Opción de pago no válida.")
		return
	}

	// Simular el pago
	if err := pago.RealizarPago(opcionDePago); err != nil {
		fmt.Printf(" [❌] Transacción fallida: %v\n", err)
	}
}

func menuReproduccion(u *Usuario) {
	if len(u.verPerfiles()) == 0 {
		fmt.Println(" [⚠️] Debes crear al menos un perfil antes de acceder al contenido.")
		return
	}

	fmt.Println("\n--- SELECCIÓN DE PERFIL ---")
	for id, p := range u.perfiles {
		fmt.Printf("%d. %s (%s)\n", id, p.nombre, p.avatar)
	}

	var perfilID int
	fmt.Print("Elige el ID del perfil: ")
	fmt.Scan(&perfilID)

	perfilElegido, err := u.getPerfilPorID(perfilID)
	if err != nil {
		fmt.Printf(" [❌] Error: %v\n", err)
		return
	}

	menuContenido(u, perfilElegido)
}

func menuContenido(u *Usuario, p *Perfil) {
	var opcion int
	for {
		fmt.Printf("\n--- PERFIL: %s ---", p.nombre)
		fmt.Println("\n1. Ver Catálogo (Listar Contenido)")
		fmt.Println("2. Buscar Contenido")
		fmt.Println("3. Simular Reproducción")
		fmt.Println("4. Ver Historial de Reproducción del Perfil")
		fmt.Println("5. Volver al Menú Anterior")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			listarContenido(Sistema.Gestor.listarTodo())
		case 2:
			ejecutarBusqueda()
		case 3:
			ejecutarSimularReproduccion(u, p)
		case 4:
			verHistorial(u, p)
		case 5:
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func listarContenido(contenido []Reproducible) {
	fmt.Println("\n--- CATÁLOGO DE CONTENIDO ---")
	if len(contenido) == 0 {
		fmt.Println(" [ℹ️] No hay contenido disponible.")
		return
	}
	for i, c := range contenido {
		tipo := "Pelicula"
		if _, isSerie := c.(*Serie); isSerie {
			tipo = "Serie"
		}
		fmt.Printf("  %d. [%s] %s (ID: %d, Género: %s)\n", i+1, tipo, c.GetTitulo(), c.GetID(), c.(*Contenido).genero)
	}
}

func ejecutarBusqueda() {
	var query string
	fmt.Print("Escribe el título (parcial o completo) a buscar: ")
	fmt.Scan(&query)

	resultados := Sistema.Gestor.buscarContenido(query)
	if len(resultados) == 0 {
		fmt.Println(" [❌] No se encontraron resultados.")
		return
	}
	listarContenido(resultados)
}

func ejecutarSimularReproduccion(u *Usuario, p *Perfil) {
	if !u.verificarSuscripcionActiva() {
		fmt.Printf(" [❌] No puedes reproducir. %v\n", ErrNoSuscripcionActiva)
		return
	}

	listarContenido(Sistema.Gestor.listarTodo())
	var idStr string
	fmt.Print("Introduce el ID del contenido a reproducir: ")
	fmt.Scan(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(" [❌] ID de contenido inválido.")
		return
	}

	contenido := Sistema.Gestor.GetContenidoPorID(id)
	if contenido == nil {
		fmt.Printf(" [❌] Error: %v\n", ErrContenidoNoEncontrado)
		return
	}

	// Llama al método Reproducir (polimorfismo)
	contenido.Reproducir(p)
}

func verHistorial(u *Usuario, p *Perfil) {
	historial := u.verHistorialReproduccion()
	fmt.Printf("\n--- HISTORIAL DE REPRODUCCIÓN - Perfil: %s ---\n", p.nombre)
	vistos := false
	for _, v := range historial.visualizaciones {
		if v.perfilID == p.id {
			vistos = true
			c := Sistema.Gestor.GetContenidoPorID(v.contenidoID)
			estado := "VISTO"
			if v.minuto < c.(*Contenido).duracion {
				estado = fmt.Sprintf("Progreso: %.1f min", v.minuto)
			}
			fmt.Printf(" - %s (%s) - Último Acceso: %s\n   [Estado: %s]\n",
				c.GetTitulo(), c.(*Contenido).genero, v.fecha.Format("02-Jan-2006 15:04"), estado)
		}
	}
	if !vistos {
		fmt.Println(" [ℹ️] Este perfil no tiene historial de reproducción.")
	}
}

/*
// Ejemplo de ejecución en la terminal (si fuera necesario):
$ go run main.go
*/
