package storage

import (
	"database/sql"
	"fmt"
	"streaming-system/pkg/auth"
	"streaming-system/pkg/billing"
	"streaming-system/pkg/content"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	User, Password, Host, Port, DBName string
}

type MySQLStorage struct {
	DB *sql.DB
}

func NewMySQLStorage(config DBConfig) (*MySQLStorage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLStorage{DB: db}, db.Ping()
}

// --- CRUD CONTENIDOS ---

func (s *MySQLStorage) LoadAllContent() ([]content.Contenible, error) {
	rows, err := s.DB.Query("SELECT CONTENIDO_ID, TITULO, DESCRIPCION FROM contenidos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []content.Contenible
	for rows.Next() {
		var id, t, d string
		rows.Scan(&id, &t, &d)
		res = append(res, content.NuevaPelicula(id, t, d, "General", "N/A", "N/A", 0))
	}
	return res, nil
}

func (s *MySQLStorage) SaveContent(c content.Contenible) error {
	// Corregido: GetDescripcion() para evitar el error de compilación
	query := `INSERT INTO contenidos (CONTENIDO_ID, TITULO, DESCRIPCION, CLASIFICACION_EDAD, ES_ESTRENO, PRECIO_COMPRA) VALUES (?, ?, ?, '13+', 1, 0)`
	_, err := s.DB.Exec(query, c.GetID(), c.GetTitulo(), c.GetDescripcion())
	return err
}

func (s *MySQLStorage) UpdateContentFull(id string, t string, d string) error {
	query := `UPDATE contenidos SET TITULO = ?, DESCRIPCION = ? WHERE CONTENIDO_ID = ?`
	_, err := s.DB.Exec(query, t, d, id)
	return err
}

func (s *MySQLStorage) DeleteContent(id string) error {
	_, err := s.DB.Exec("DELETE FROM contenidos WHERE CONTENIDO_ID = ?", id)
	return err
}

// --- USUARIOS, PERFILES Y SUSCRIPCIONES ---

func (s *MySQLStorage) SaveUser(u *auth.Usuario) error {
	_, err := s.DB.Exec("INSERT INTO usuarios (USUARIO_ID, NOMBRE_USUARIO, EMAIL, PASSWORD_HASH, METODOS_PAGO_METODO_PAGO_ID) VALUES (?, ?, ?, ?, 0)", u.GetID(), u.GetNombre(), u.GetCorreo(), u.GetContraseniaHash())
	return err
}

func (s *MySQLStorage) LoadUserByEmail(e string) (*auth.Usuario, error) {
	var id, n, em, h string
	err := s.DB.QueryRow("SELECT USUARIO_ID, NOMBRE_USUARIO, EMAIL, PASSWORD_HASH FROM usuarios WHERE EMAIL = ?", e).Scan(&id, &n, &em, &h)
	if err != nil {
		return nil, err
	}
	return auth.RecreateUsuarioFromDB(id, n, em, h, nil), nil
}

func (s *MySQLStorage) SaveProfile(uid string, p *auth.Perfil) error {
	_, err := s.DB.Exec("INSERT INTO perfiles (PERFIL_ID, USUARIO_ID, NOMBRE_PERFIL, CLASIFICACION_EDAD_MAXIMA, USUARIOS_USUARIO_ID) VALUES (?, ?, ?, 18, ?)", p.GetID(), uid, p.GetNombre(), uid)
	return err
}

func (s *MySQLStorage) LoadProfilesByUserID(uid string) ([]*auth.Perfil, error) {
	rows, _ := s.DB.Query("SELECT PERFIL_ID, NOMBRE_PERFIL FROM perfiles WHERE USUARIOS_USUARIO_ID = ?", uid)
	var res []*auth.Perfil
	for rows.Next() {
		var id int
		var n string
		rows.Scan(&id, &n)
		res = append(res, auth.NuevoPerfil(id, n))
	}
	return res, nil
}

func (s *MySQLStorage) SaveSubscription(uid string, sub *billing.Suscripcion) error {
	p := sub.GetPlan()
	_, err := s.DB.Exec("INSERT INTO planes_suscripcion (USUARIO_ID, METODO_PAGO_ID, TIPO_PLAN, FECHA_INICIO, FECHA_FIN, ESTADO_SUSCRIPCION, PRECIO_MENSUAL_ANUAL, FECHA_PROXIMO_PAGO, METODOS_PAGO_METODO_PAGO_ID, USUARIOS_USUARIO_ID) VALUES (?, 1, ?, ?, ?, ?, ?, ?, 1, ?)", uid, p.GetID(), sub.GetFechaInicio(), sub.GetFechaFin(), sub.GetEstado(), p.GetPrecio(), sub.GetFechaFin(), uid)
	return err
}

func (s *MySQLStorage) LoadActiveSubscription(uid string) (*billing.Suscripcion, error) {
	var tp, es string
	var si, sf time.Time
	var pr float64
	err := s.DB.QueryRow("SELECT TIPO_PLAN, FECHA_INICIO, FECHA_FIN, ESTADO_SUSCRIPCION, PRECIO_MENSUAL_ANUAL FROM planes_suscripcion WHERE USUARIOS_USUARIO_ID = ? AND ESTADO_SUSCRIPCION = 'ACTIVO' LIMIT 1", uid).Scan(&tp, &si, &sf, &es, &pr)
	if err != nil {
		return nil, err
	}
	return billing.RecreateSuscripcionFromDB("S-"+uid, billing.NuevoPlan(tp, tp, pr, pr), si, sf, es), nil
}
