package auth

import (
	"streaming-system/pkg/billing"
)

type Usuario struct {
	id, nombre, correo, contraseniaHash string
	perfiles                            []*Perfil
	suscripcion                         *billing.Suscripcion
	autenticacion                       *Autenticacion
}

func NuevoUsuario(id, nombre, correo, pass string) *Usuario {
	return &Usuario{
		id: id, nombre: nombre, correo: correo,
		contraseniaHash: "HASH_" + pass,
		perfiles:        make([]*Perfil, 0),
		autenticacion:   &Autenticacion{},
	}
}

func RecreateUsuarioFromDB(id, nombre, correo, hash string, sus *billing.Suscripcion) *Usuario {
	return &Usuario{
		id: id, nombre: nombre, correo: correo, contraseniaHash: hash,
		perfiles: make([]*Perfil, 0), autenticacion: &Autenticacion{},
		suscripcion: sus,
	}
}

// Getters y Métodos de Negocio
func (u *Usuario) GetID() string                        { return u.id }
func (u *Usuario) GetNombre() string                    { return u.nombre }
func (u *Usuario) GetCorreo() string                    { return u.correo }
func (u *Usuario) GetContraseniaHash() string           { return u.contraseniaHash }
func (u *Usuario) GetPerfiles() []*Perfil               { return u.perfiles }
func (u *Usuario) GetSuscripcion() *billing.Suscripcion { return u.suscripcion }

func (u *Usuario) IniciarSesion(correo, pass string) error {
	if u.correo == correo && u.contraseniaHash == "HASH_"+pass {
		u.autenticacion.GenerarToken(u.id)
		return nil
	}
	return ErrCredencialesInvalidas
}

func (u *Usuario) AsignarSuscripcion(s *billing.Suscripcion) { u.suscripcion = s }
func (u *Usuario) TieneSuscripcionActiva() bool {
	return u.suscripcion != nil && u.suscripcion.EstaActiva()
}

func (u *Usuario) CrearPerfil(nombre string) *Perfil {
	p := NuevoPerfil(len(u.perfiles)+1, nombre)
	u.perfiles = append(u.perfiles, p)
	return p
}

func (u *Usuario) AgregarPerfilExistente(p *Perfil) { u.perfiles = append(u.perfiles, p) }
