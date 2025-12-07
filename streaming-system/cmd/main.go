package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"streaming-system/pkg/auth"
	"streaming-system/pkg/billing"
	"streaming-system/pkg/content"
	"strings"
)

// Globales Mock
var gestor *content.GestorDeContenido
var usuarios map[string]*auth.Usuario // Clave: Correo
var planesDisponibles map[string]*billing.Plan
var usuarioActual *auth.Usuario
var perfilActual *auth.Perfil

// Constante para simular el correo del administrador (Control de Acceso)
const CorreoAdmin = "admin@stream.com"

func main() {
	fmt.Println("=========================================")
	fmt.Println("🎬 Sistema de Gestión de Streaming (Go)")
	fmt.Println("=========================================")

	inicializarDatosMock()
	menuPrincipal()
}

// Función para inicializar datos de prueba.
func inicializarDatosMock() {
	usuarios = make(map[string]*auth.Usuario)
	gestor = content.NuevoGestorDeContenido()

	// 1. Contenido Mock
	peli1 := content.NuevaPelicula("C001", "El Origen de los Go-Fers", "Ciencia ficción sobre goroutines.", "Ciencia Ficción", "Gopher Nolan", "LinkTrailer", 150.5)
	serie1 := content.NuevaSerie("C002", "Punteros Peligrosos", "Thriller sobre desreferenciación.", "Thriller", 2)
	serie1.AgregarEpisodio("El Secreto del Nil", 45.0)
	serie1.AgregarEpisodio("La Fuga de la Interfaz", 48.0)

	gestor.InsertarContenido(peli1)
	gestor.InsertarContenido(serie1)

	// 2. Planes Mock
	planesDisponibles = map[string]*billing.Plan{
		"1": billing.NuevoPlan("P01", "Básico", 9.99, 9.99),
		"2": billing.NuevoPlan("P02", "Premium", 19.99, 19.99),
	}

	// 3. Crear el usuario administrador para pruebas
	adminUser := auth.NuevoUsuario("U0", "Admin Master", CorreoAdmin, "secure_admin_pass")
	usuarios[CorreoAdmin] = adminUser
}

// LeerEntrada lee una línea del terminal.
func leerEntrada() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// EsAdmin verifica si el usuario logueado es el administrador.
func EsAdmin(u *auth.Usuario) bool {
	return u != nil && u.GetID() == "U0"
}

// Menu Principal (MODIFICADO: Opción de Administración en menú desconectado)
func menuPrincipal() {
	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		if usuarioActual == nil {
			fmt.Println("1. Registrar nuevo usuario")
			fmt.Println("2. Iniciar Sesión")
			fmt.Println("3. Salir")
			fmt.Println("A. Módulo de Administración (Requiere Admin Login)") // OPCIÓN VISIBLE DESCONECTADO
		} else {
			fmt.Printf("👤 Conectado como: %s | Perfil: %s\n", usuarioActual.GetID(), func() string {
				if perfilActual != nil {
					return perfilActual.GetNombre()
				}
				return "N/A"
			}())

			// OPCIONES DE USUARIO ESTÁNDAR
			fmt.Println("4. Ver Planes y Suscribirse/Pagar")
			fmt.Println("5. Gestionar Perfiles")
			fmt.Println("6. Ver Contenido Disponible")
			fmt.Println("7. Ver Historial de Reproducción (Perfil)")
			fmt.Println("9. Cerrar Sesión")
		}

		fmt.Print("Elige una opción: ")
		opcion := strings.ToUpper(leerEntrada()) // Convertir a mayúsculas para manejar 'A'

		if usuarioActual == nil {
			switch opcion {
			case "1":
				simularRegistro()
			case "2":
				simularInicioSesion()
			case "3":
				fmt.Println("👋 ¡Gracias por usar el sistema!")
				return
			case "A": // CASO AHORA ACCESIBLE SIN INICIAR SESIÓN
				// Lógica para forzar la autenticación del administrador antes de entrar al módulo
				fmt.Println("\n--- ACCESO DE ADMINISTRADOR ---")
				fmt.Print("Correo Admin: ")
				correo := leerEntrada()
				fmt.Print("Contraseña Admin: ")
				contrasenia := leerEntrada()

				user, ok := usuarios[correo]
				if !ok || user.GetID() != "U0" {
					fmt.Println("❌ Acceso denegado: Credenciales o privilegios inválidos.")
					return
				}
				err := user.IniciarSesion(correo, contrasenia)
				if err != nil {
					fmt.Println("❌ Acceso denegado:", err)
					return
				}
				// Si la autenticación es exitosa, se abre el menú de administración
				fmt.Println("✅ Autenticación de Administrador Exitosa.")
				menuAdministracion()
				// Tras salir de menuAdministracion(), el usuario debe cerrar sesión
				user.CerrarSesion()

			default:
				fmt.Println("Opción no válida.")
			}
		} else {
			switch opcion {
			case "4":
				simularGestionSuscripcion()
			case "5":
				simularGestionPerfiles()
			case "6":
				verContenidoDisponible()
			case "7":
				verHistorialReproduccion()
			case "9":
				usuarioActual.CerrarSesion()
				usuarioActual = nil
				perfilActual = nil
				fmt.Println("✅ Sesión cerrada correctamente.")
			default:
				fmt.Println("Opción no válida.")
			}
		}
	}
}

// 1. Módulo A: Registro y Autenticación
func simularRegistro() {
	fmt.Println("\n--- REGISTRO DE USUARIO ---")
	fmt.Print("Correo (ej. user@test.com): ")
	correo := leerEntrada()
	fmt.Print("Nombre: ")
	nombre := leerEntrada()
	fmt.Print("Contraseña: ")
	contrasenia := leerEntrada()

	if _, ok := usuarios[correo]; ok {
		fmt.Println("❌ Error: Ya existe un usuario con ese correo.")
		return
	}

	// Si el correo es el de admin, se le asigna el ID fijo U0, sino se asigna un nuevo ID incremental
	var id string
	if correo == CorreoAdmin {
		id = "U0"
	} else {
		id = fmt.Sprintf("U%d", len(usuarios)+1)
	}

	nuevoUsuario := auth.NuevoUsuario(id, nombre, correo, contrasenia)
	usuarios[correo] = nuevoUsuario
	fmt.Println("✅ Usuario registrado con éxito. Ahora puede iniciar sesión.")
}

func simularInicioSesion() {
	fmt.Println("\n--- INICIO DE SESIÓN ---")
	fmt.Print("Correo: ")
	correo := leerEntrada()
	fmt.Print("Contraseña: ")
	contrasenia := leerEntrada()

	user, ok := usuarios[correo]
	if !ok {
		fmt.Println("❌ Error:", auth.ErrCredencialesInvalidas)
		return
	}

	err := user.IniciarSesion(correo, contrasenia)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	usuarioActual = user
}

// 2. Módulo C: Suscripción y Pago
func simularGestionSuscripcion() {
	if usuarioActual.TieneSuscripcionActiva() {
		fmt.Printf("ℹ️ Ya tiene una suscripción **%s** activa. ¿Desea renovarla o cancelarla? (R/C/N): ", usuarioActual.GetSuscripcion().GetPlan().GetNombre())
		op := strings.ToUpper(leerEntrada())
		if op == "C" {
			usuarioActual.GetSuscripcion().Cancelar()
			fmt.Println("✅ Suscripción cancelada.")
			return
		} else if op != "R" {
			return
		}
	}

	fmt.Println("\n--- PLANES DISPONIBLES ---")
	for id, p := range planesDisponibles {
		fmt.Printf("%s. %s ($%.2f/mes)\n", id, p.GetNombre(), p.GetPrecio())
	}

	fmt.Print("Elige un Plan (ID): ")
	planID := leerEntrada()
	planElegido, ok := planesDisponibles[planID]
	if !ok {
		fmt.Println("❌ Plan no válido.")
		return
	}

	// Simulación de Pago
	suscripcion := billing.NuevaSuscripcion(fmt.Sprintf("S%s", usuarioActual.GetID()), planElegido)
	pago := billing.NuevoPago(fmt.Sprintf("T%s", usuarioActual.GetID()), planElegido.GetPrecio(), billing.TarjetaCredito)

	fmt.Println("\n--- SIMULACIÓN DE PAGO ---")
	fmt.Print("Elige Opción de Pago (1: Tarjeta, 2: PayPal): ")
	opcionPago := leerEntrada()
	var pagoOpcion billing.OpcionPago
	if opcionPago == "1" {
		pagoOpcion = billing.TarjetaCredito
	} else if opcionPago == "2" {
		pagoOpcion = billing.PayPal
	} else {
		fmt.Println("❌ Opción de pago no válida. Usando Tarjeta por defecto.")
		pagoOpcion = billing.TarjetaCredito
	}

	err := pago.RegistrarPago(planElegido.GetPrecio(), pagoOpcion, suscripcion)
	if err != nil {
		fmt.Println("❌ Fallo en el pago:", err)
		return
	}

	// Asignar la suscripción al usuario solo si el pago fue exitoso
	usuarioActual.AsignarSuscripcion(suscripcion)
	fmt.Println("✅ Suscripción adquirida con éxito.")
}

// 3. Módulo A: Gestión de Perfiles
func simularGestionPerfiles() {
	fmt.Println("\n--- GESTIÓN DE PERFILES ---")
	perfiles := usuarioActual.GetPerfiles()

	if len(perfiles) == 0 {
		fmt.Println("ℹ️ No tiene perfiles creados.")
	} else {
		fmt.Println("Perfiles existentes:")
		for _, p := range perfiles {
			fmt.Printf("ID %d: %s\n", p.GetID(), p.GetNombre())
		}
	}

	fmt.Print("Elige una acción (C: Crear, E: Elegir Perfil, N: Nada): ")
	accion := strings.ToUpper(leerEntrada())

	if accion == "C" {
		fmt.Print("Nombre del nuevo perfil: ")
		nombre := leerEntrada()
		nuevo := usuarioActual.CrearPerfil(nombre)
		fmt.Printf("✅ Perfil '%s' creado con ID %d.\n", nuevo.GetNombre(), nuevo.GetID())
		perfilActual = nuevo // Asignar automáticamente el nuevo perfil
	} else if accion == "E" {
		fmt.Print("ID del perfil a elegir: ")
		idStr := leerEntrada()
		id, _ := strconv.Atoi(idStr)
		for _, p := range perfiles {
			if p.GetID() == id {
				perfilActual = p
				fmt.Printf("✅ Perfil activo cambiado a: %s\n", p.GetNombre())
				return
			}
		}
		fmt.Println("❌ ID de perfil no encontrado.")
	}
}

// 4. Módulo B: Ver Contenido (Usado para ver el catálogo)
func verContenidoDisponible() {
	fmt.Println("\n--- CONTENIDO DISPONIBLE ---")
	// La restricción de suscripción solo se aplica si el usuario no es admin
	if usuarioActual == nil || (!EsAdmin(usuarioActual) && !usuarioActual.TieneSuscripcionActiva()) {
		fmt.Println("❌ **AUTORIZACIÓN REQUERIDA**: Necesita una suscripción activa para ver el contenido.")
		return
	}

	i := 1
	for _, c := range gestor.ObtenerTodo() {
		switch v := c.(type) {
		case *content.Pelicula:
			fmt.Printf("%d. [Película] ID: %s | Título: **%s** | Género: %s\n", i, v.GetID(), v.GetTitulo(), v.GetGenero())
		case *content.Serie:
			fmt.Printf("%d. [Serie] ID: %s | Título: **%s** | Género: %s | Eps: %d\n", i, v.GetID(), v.GetTitulo(), v.GetGenero(), len(v.ObtenerEpisodios()))
		}
		i++
	}
}

// 5. Módulo de Administración
func menuAdministracion() {
	for {
		fmt.Println("\n--- MÓDULO DE ADMINISTRACIÓN (CRUD) ---")
		fmt.Println("1. Ver Contenido (Leer)")
		fmt.Println("2. Agregar Nuevo Contenido (Crear)")
		fmt.Println("3. Modificar Título (Actualizar)")
		fmt.Println("4. Eliminar Contenido (Eliminar)")
		fmt.Println("5. Volver al Menú Principal")

		fmt.Print("Elige una opción: ")
		opcion := leerEntrada()

		switch opcion {
		case "1":
			// No necesita suscripción para ver el catálogo desde el panel de admin
			verContenidoDisponible()
		case "2":
			agregarContenido()
		case "3":
			modificarContenido()
		case "4":
			eliminarContenido()
		case "5":
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func agregarContenido() {
	fmt.Println("\n--- AGREGAR CONTENIDO ---")
	fmt.Print("Tipo (P: Película, S: Serie): ")
	tipo := strings.ToUpper(leerEntrada())
	fmt.Print("ID (ej. C003): ")
	id := leerEntrada()
	fmt.Print("Título: ")
	titulo := leerEntrada()
	fmt.Print("Descripción: ")
	descripcion := leerEntrada()
	fmt.Print("Género: ")
	genero := leerEntrada()

	if tipo == "P" {
		// Película
		fmt.Print("Director: ")
		director := leerEntrada()
		fmt.Print("Duración (min): ")
		duracionStr := leerEntrada()

		// --- CORRECCIÓN FLOAT64 A FLOAT32 ---
		duracion64, _ := strconv.ParseFloat(duracionStr, 64)
		duracion32 := float32(duracion64) // Conversión explícita

		nuevaPeli := content.NuevaPelicula(id, titulo, descripcion, genero, director, "N/A", duracion32)
		gestor.InsertarContenido(nuevaPeli)
		fmt.Printf("✅ Película '%s' agregada.\n", titulo)
	} else if tipo == "S" {
		// Serie
		fmt.Print("Temporadas: ")
		temporadasStr := leerEntrada()
		temporadas, _ := strconv.Atoi(temporadasStr)

		nuevaSerie := content.NuevaSerie(id, titulo, descripcion, genero, temporadas)
		gestor.InsertarContenido(nuevaSerie)
		fmt.Printf("✅ Serie '%s' agregada. Añade episodios a través de la interfaz de desarrollo.\n", titulo)
	} else {
		fmt.Println("❌ Tipo de contenido no reconocido.")
	}
}

func modificarContenido() {
	fmt.Println("\n--- MODIFICAR TÍTULO ---")
	verContenidoDisponible()
	fmt.Print("ID del Contenido a modificar: ")
	id := leerEntrada()

	fmt.Print("Nuevo Título: ")
	nuevoTitulo := leerEntrada()

	err := gestor.ActualizarContenidoMetadata(id, nuevoTitulo)
	if err != nil {
		fmt.Println("❌ Error al modificar:", err)
	} else {
		fmt.Printf("✅ Título del contenido %s actualizado a '%s'.\n", id, nuevoTitulo)
	}
}

func eliminarContenido() {
	fmt.Println("\n--- ELIMINAR CONTENIDO ---")
	verContenidoDisponible()
	fmt.Print("ID del Contenido a eliminar: ")
	id := leerEntrada()

	gestor.BorrarContenido(id)
	fmt.Printf("✅ Contenido %s eliminado del catálogo.\n", id)
}

// 6. Módulo D: Ver Historial (Ahora es la Opción 7)
func verHistorialReproduccion() {
	if perfilActual == nil {
		fmt.Println("⚠️ Debe elegir un perfil activo (Opción 5) para ver su historial.")
		return
	}

	historial := usuarioActual.GetHistorialReproduccion()
	if len(historial.GetVisualizaciones()) == 0 {
		fmt.Printf("\nℹ️ El perfil **%s** no tiene historial de reproducción.\n", perfilActual.GetNombre())
		return
	}

	fmt.Printf("\n--- HISTORIAL DE REPRODUCCIÓN - PERFIL: %s ---\n", perfilActual.GetNombre())
	for _, v := range historial.GetVisualizaciones() {
		titulo := v.GetContenido().GetTitulo()
		progreso := ""

		// Usando Getters para cumplir encapsulación
		if v.GetGuardarProgreso() > 0 {
			progreso = fmt.Sprintf(" (Progreso: min %d)", v.GetGuardarProgreso())
		}

		// Usando Getters para cumplir encapsulación
		fmt.Printf("• %s | %s%s\n", titulo, v.GetFecha().Format("02 Jan 15:04"), progreso)
	}
}
