package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"streaming-system/pkg/auth"
	"streaming-system/pkg/billing"
	"streaming-system/pkg/content"
	"streaming-system/pkg/playback"
	"strings"
)

// Globales Mock
var gestor *content.GestorDeContenido
var usuarios map[string]*auth.Usuario // Clave: Correo
var planesDisponibles map[string]*billing.Plan
var usuarioActual *auth.Usuario
var perfilActual *auth.Perfil

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
}

// LeerEntrada lee una línea del terminal.
func leerEntrada() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Menu Principal
func menuPrincipal() {
	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		if usuarioActual == nil {
			fmt.Println("1. Registrar nuevo usuario")
			fmt.Println("2. Iniciar Sesión")
			fmt.Println("3. Salir")
		} else {
			fmt.Printf("👤 Conectado como: %s | Perfil: %s\n", usuarioActual.GetID(), func() string {
				if perfilActual != nil {
					return perfilActual.GetNombre()
				}
				return "N/A"
			}())
			fmt.Println("4. Ver Planes y Suscribirse/Pagar")
			fmt.Println("5. Gestionar Perfiles")
			fmt.Println("6. Ver Contenido Disponible")
			fmt.Println("7. Simular Reproducción")
			fmt.Println("8. Ver Historial de Reproducción (Perfil)")
			fmt.Println("9. Cerrar Sesión")
		}

		fmt.Print("Elige una opción: ")
		opcion := leerEntrada()

		if usuarioActual == nil {
			switch opcion {
			case "1":
				simularRegistro()
			case "2":
				simularInicioSesion()
			case "3":
				fmt.Println("👋 ¡Gracias por usar el sistema!")
				return
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
				simularReproduccion()
			case "8":
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

	nuevoUsuario := auth.NuevoUsuario(fmt.Sprintf("U%d", len(usuarios)+1), nombre, correo, contrasenia)
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

// 4. Módulo B: Ver Contenido
func verContenidoDisponible() {
	fmt.Println("\n--- CONTENIDO DISPONIBLE ---")
	if !usuarioActual.TieneSuscripcionActiva() {
		fmt.Println("❌ **AUTORIZACIÓN REQUERIDA**: Necesita una suscripción activa para ver el contenido.")
		return
	}

	if perfilActual == nil {
		fmt.Println("⚠️ Necesita elegir un perfil activo (Opción 5) antes de ver contenido.")
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

// 5. Módulo D: Simular Reproducción y Progreso
func simularReproduccion() {
	if !usuarioActual.TieneSuscripcionActiva() {
		fmt.Println("❌ **AUTORIZACIÓN REQUERIDA**: Necesita una suscripción activa para reproducir.")
		return
	}
	if perfilActual == nil {
		fmt.Println("⚠️ Necesita elegir un perfil activo (Opción 5) antes de reproducir contenido.")
		return
	}
	verContenidoDisponible() // Mostrar el listado

	fmt.Print("\nID del Contenido a reproducir (ej. C001, C002): ")
	contenidoID := leerEntrada()

	c, err := gestor.BuscarContenido(contenidoID)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	// 1. Iniciar/Reanudar Reproducción
	historial := usuarioActual.GetHistorialReproduccion()
	vis := historial.ObtenerUltimaVisualizacionPorContenido(contenidoID)

	if vis == nil {
		// Nuevo inicio de visualización
		vis = playback.NuevoVisualizacion(fmt.Sprintf("V%s-%d", perfilActual.GetID(), len(historial.GetVisualizaciones())+1), c)
		vis.GetContenido().Reproducir()
	} else {
		// Reanudar visualización
		vis.ReproducirDesdePunto()
	}

	// 2. Simular guardar progreso
	fmt.Print("Simular interrupción. ¿Guardar progreso en el minuto? (0-100): ")
	progresoStr := leerEntrada()
	progreso, _ := strconv.Atoi(progresoStr)

	if progreso > 0 {
		// CORRECCIÓN: Usar el método público GuardarProgreso()
		vis.GuardarProgreso(progreso)
		historial.AgregarVisualizacion(vis)
		fmt.Println("✅ Visualización registrada en el historial del perfil", perfilActual.GetNombre())
	} else {
		fmt.Println("⏭️ No se guardó progreso. Reproducción terminada/descartada.")
	}
}

// 6. Módulo D: Ver Historial
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

		// CORRECCIÓN: Usar el método público GetGuardarProgreso()
		if v.GetGuardarProgreso() > 0 {
			// CORRECCIÓN: Usar el método público GetGuardarProgreso()
			progreso = fmt.Sprintf(" (Progreso: min %d)", v.GetGuardarProgreso())
		}

		// CORRECCIÓN: Usar el método público GetFecha()
		fmt.Printf("• %s | %s%s\n", titulo, v.GetFecha().Format("02 Jan 15:04"), progreso)
	}
}
