package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"streaming-system/pkg/auth"
	"streaming-system/pkg/billing"
	"streaming-system/pkg/content"
	"streaming-system/pkg/storage"
	"time"
)

var (
	gestor            *content.GestorDeContenido
	dbStore           *storage.MySQLStorage
	planesDisponibles map[string]*billing.Plan
	usuarioActual     *auth.Usuario
	perfilActual      *auth.Perfil
	tmpl              *template.Template
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("🎬 Servidor Web StreamGo - FULL SYSTEM")
	fmt.Println("=========================================")

	s, err := storage.NewMySQLStorage(storage.DBConfig{
		User: "root", Password: "Kevin1994Alex", Host: "localhost", Port: "3306", DBName: "BDD_Streaming",
	})
	if err != nil {
		fmt.Printf("❌ DB Error: %v\n", err)
		os.Exit(1)
	}
	dbStore = s
	defer s.DB.Close()

	inicializarDatos()
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	// RUTAS
	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/profiles", handleProfiles)
	http.HandleFunc("/profiles/create", handleCreateProfile)
	http.HandleFunc("/profiles/select", handleSelectProfile)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/checkout", handleCheckout)
	http.HandleFunc("/checkout/buy", handleBuyPlan)
	http.HandleFunc("/admin", handleAdmin)
	http.HandleFunc("/admin/add", handleAdminAdd)
	http.HandleFunc("/admin/update", handleAdminUpdate)
	http.HandleFunc("/admin/delete", handleAdminDelete)
	http.HandleFunc("/logout", handleLogout)

	port := ":8080"
	fmt.Printf("🚀 Servidor corriendo en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func inicializarDatos() {
	planesDisponibles = map[string]*billing.Plan{
		"1": billing.NuevoPlan("P01", "Plan Básico", 9.99, 5.0),
		"2": billing.NuevoPlan("P02", "Plan Premium", 19.99, 10.0),
	}
	lista, _ := dbStore.LoadAllContent()
	gestor = content.NuevoGestorDeContenido(lista)
}

// --- MANEJADORES ---

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("password")
		u, _ := dbStore.LoadUserByEmail(email)
		if u != nil && u.IniciarSesion(email, pass) == nil {
			usuarioActual = u
			sub, _ := dbStore.LoadActiveSubscription(u.GetID())
			if sub != nil {
				usuarioActual.AsignarSuscripcion(sub)
			}
			if email == "admin@stream.com" {
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, "/profiles", http.StatusSeeOther)
			return
		}
	}
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		u := auth.NuevoUsuario(fmt.Sprintf("%d", time.Now().Unix()), r.FormValue("nombre"), r.FormValue("email"), r.FormValue("password"))
		dbStore.SaveUser(u)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func handleProfiles(w http.ResponseWriter, r *http.Request) {
	if usuarioActual == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	p, _ := dbStore.LoadProfilesByUserID(usuarioActual.GetID())
	tmpl.ExecuteTemplate(w, "profiles.html", struct{ Perfiles []*auth.Perfil }{p})
}

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	if usuarioActual == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	dbStore.SaveProfile(usuarioActual.GetID(), auth.NuevoPerfil(int(time.Now().Unix()%1000), r.FormValue("nombre")))
	http.Redirect(w, r, "/profiles", http.StatusSeeOther)
}

func handleSelectProfile(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	perfiles, _ := dbStore.LoadProfilesByUserID(usuarioActual.GetID())
	for _, p := range perfiles {
		if p.GetID() == id {
			perfilActual = p
			break
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	if usuarioActual == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if usuarioActual.GetSuscripcion() == nil {
		http.Redirect(w, r, "/checkout", http.StatusSeeOther)
		return
	}
	if perfilActual == nil {
		http.Redirect(w, r, "/profiles", http.StatusSeeOther)
		return
	}

	mapa := gestor.ObtenerTodo()
	var lista []content.Contenible
	for _, c := range mapa {
		lista = append(lista, c)
	}
	tmpl.ExecuteTemplate(w, "dashboard.html", struct {
		Usuario    *auth.Usuario
		Perfil     *auth.Perfil
		Contenidos []content.Contenible
	}{usuarioActual, perfilActual, lista})
}

func handleCheckout(w http.ResponseWriter, r *http.Request) {
	if usuarioActual == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "checkout.html", nil)
}

func handleBuyPlan(w http.ResponseWriter, r *http.Request) {
	p := planesDisponibles[r.FormValue("plan_id")]
	s := billing.NuevaSuscripcion("S-"+usuarioActual.GetID(), p)
	dbStore.SaveSubscription(usuarioActual.GetID(), s)
	usuarioActual.AsignarSuscripcion(s)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	if usuarioActual == nil || usuarioActual.GetCorreo() != "admin@stream.com" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	mapa := gestor.ObtenerTodo()
	var lista []content.Contenible
	for _, c := range mapa {
		lista = append(lista, c)
	}
	tmpl.ExecuteTemplate(w, "admin.html", lista)
}

func handleAdminAdd(w http.ResponseWriter, r *http.Request) {
	n := content.NuevaPelicula(r.FormValue("id"), r.FormValue("titulo"), r.FormValue("descripcion"), "General", "N/A", "N/A", 0)
	gestor.InsertarContenido(n)
	dbStore.SaveContent(n)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func handleAdminUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	t := r.FormValue("titulo")
	d := r.FormValue("descripcion")
	if err := gestor.ActualizarContenidoMetadata(id, t); err == nil {
		dbStore.UpdateContentFull(id, t, d)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func handleAdminDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	gestor.BorrarContenido(id)
	dbStore.DeleteContent(id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	usuarioActual = nil
	perfilActual = nil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
