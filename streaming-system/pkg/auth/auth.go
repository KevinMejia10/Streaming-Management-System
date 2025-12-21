package auth

import (
	"errors"
	"time"
)

// ErrCredencialesInvalidas es un error personalizado para fallos de inicio de sesión.
// ¡Declaración única para todo el paquete auth!
var ErrCredencialesInvalidas = errors.New("credenciales inválidas")

// Autenticacion representa la sesión de un usuario.
type Autenticacion struct {
	token      string    // token de sesión
	expireDate time.Time // fecha de expiración del token
}

// GenerarToken simula la generación de un nuevo token de sesión.
func (a *Autenticacion) GenerarToken(usuarioID string) string {
	// Simulación simple: un token es el ID del usuario más un timestamp
	a.token = usuarioID + "_" + time.Now().Format("20060102150405")
	a.expireDate = time.Now().Add(24 * time.Hour)
	return a.token
}

// VerificarToken simula la verificación de la validez de un token.
func (a *Autenticacion) VerificarToken() bool {
	return time.Now().Before(a.expireDate) && a.token != ""
}

// RenovarToken simula la extensión de la vida del token.
func (a *Autenticacion) RenovarToken() {
	a.expireDate = time.Now().Add(48 * time.Hour)
}

// RevocarToken simula el cierre de sesión, invalidando el token.
func (a *Autenticacion) RevocarToken() {
	a.token = ""
	a.expireDate = time.Time{}
}

// GetToken exporta el valor del token.
func (a *Autenticacion) GetToken() string {
	return a.token
}
