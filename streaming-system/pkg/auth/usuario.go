package auth

import (
	"fmt"
	"streaming-system/pkg/billing"
	"streaming-system/pkg/playback"
)

// Usuario representa un usuario del sistema de streaming.
type Usuario struct {
	id                    string // ID único del usuario
	nombre                string
	contraseniaHash       string // Contraseña hasheada (simulada)
	correo                string
	perfiles              []*Perfil
	suscripcion           *billing.Suscripcion
	autenticacion         *Autenticacion
	historialReproduccion *playback.HistorialReproduccion
}

// NuevoUsuario es el constructor para Usuario.
func NuevoUsuario(id, nombre, correo, contrasenia string) *Usuario {
	return &Usuario{
		id:                    id,
		nombre:                nombre,
		correo:                correo,
		contraseniaHash:       hashPassword(contrasenia), // Simulación de hashing
		perfiles:              make([]*Perfil, 0),
		autenticacion:         &Autenticacion{},
		historialReproduccion: &playback.HistorialReproduccion{},
	}
}

// hashPassword simula el proceso de hashing de contraseñas.
func hashPassword(contrasenia string) string {

	return "HASH_" + contrasenia
}

// GetID devuelve el ID del usuario.
func (u *Usuario) GetID() string {
	return u.id
}

// GetPerfiles devuelve la lista de perfiles del usuario.
func (u *Usuario) GetPerfiles() []*Perfil {
	return u.perfiles
}

// GetAutenticacion devuelve la instancia de Autenticacion del usuario.
func (u *Usuario) GetAutenticacion() *Autenticacion {
	return u.autenticacion
}

// GetSuscripcion devuelve la suscripción del usuario.
func (u *Usuario) GetSuscripcion() *billing.Suscripcion {
	return u.suscripcion
}

// GetHistorialReproduccion devuelve el historial del usuario.
func (u *Usuario) GetHistorialReproduccion() *playback.HistorialReproduccion {
	return u.historialReproduccion
}

// IniciarSesion simula la verificación de credenciales y la generación del token.
func (u *Usuario) IniciarSesion(correo string, contrasenia string) error {
	if u.correo != correo || u.contraseniaHash != hashPassword(contrasenia) {
		return ErrCredencialesInvalidas
	}
	u.autenticacion.GenerarToken(u.id)
	fmt.Printf("✅ Sesión iniciada. Token: %s\n", u.autenticacion.GetToken())
	return nil
}

// CerrarSesion revoca el token de autenticación.
func (u *Usuario) CerrarSesion() {
	u.autenticacion.RevocarToken()
}

// CrearPerfil simula la creación y asignación de un nuevo perfil.
func (u *Usuario) CrearPerfil(nombre string) *Perfil {
	// Generar un ID simple para el nuevo perfil
	perfilID := len(u.perfiles) + 1
	nuevo := NuevoPerfil(perfilID, nombre)
	u.perfiles = append(u.perfiles, nuevo)
	return nuevo
}

// EliminarPerfilPorID simula la eliminación de un perfil.
func (u *Usuario) EliminarPerfilPorID(perfilID int) {
	for i, p := range u.perfiles {
		if p.id == perfilID {
			u.perfiles = append(u.perfiles[:i], u.perfiles[i+1:]...)
			return
		}
	}
}

// AsignarSuscripcion establece la suscripción del usuario.
func (u *Usuario) AsignarSuscripcion(s *billing.Suscripcion) {
	u.suscripcion = s
}

// TieneSuscripcionActiva verifica si el usuario tiene una suscripción válida.
func (u *Usuario) TieneSuscripcionActiva() bool {
	if u.suscripcion == nil {
		return false
	}
	return u.suscripcion.EstaActiva()
}
