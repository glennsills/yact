package ui

import (
	"html/template"
	"log"
	"net/http"

	"github.com/glennsills/yact/config"
	userData "github.com/glennsills/yact/db/usersData"
	userService "github.com/glennsills/yact/services/userSvc"
	"github.com/glennsills/yact/services/userSvc/validation"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var templates = template.Must(template.ParseFiles(
	"ui/views/default.html",
	"ui/views/users/listusers.html",
	"ui/views/users/edituser.html",
	"ui/views/users/newuser.html",
	"ui/views/users/deleteuser.html",
	"ui/views/users/userEmail.html"))

//var validPath = regexp.MustCompile("^/(edit|save|view|users)/([a-zA-Z0-9]+)$")

func initUserService() (*userService.UserService, error) {
	db, err := gorm.Open(postgres.Open(config.App.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	repo := userData.NewUserRepository(db)
	service := userService.NewUserService(repo)
	return service, err
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "default.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	service, err := initUserService()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	list, err := service.ListActiveUsers()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	err = templates.ExecuteTemplate(w, "listusers.html", list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := templates.ExecuteTemplate(w, "edituser.html", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "newuser.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "edituser.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//do some stuff
	err := templates.ExecuteTemplate(w, "deleteuser.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func validateEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	isValid := validation.IsValidEmail(email)

	err := templates.ExecuteTemplate(w, "userEmail.html", isValid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AssignHandlers() *http.ServeMux {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("GET /users/", listUserHandler)
	mux.HandleFunc("GET /user/{id}", editUserHandler)
	mux.HandleFunc("DELETE /user/{id}", deleteUserHandler)
	mux.HandleFunc("POST /user", newUserHandler)
	mux.HandleFunc("PUT /user", updateUserHandler)
	mux.HandleFunc("POST /validate/email", validateEmailHandler)

	return mux
}
