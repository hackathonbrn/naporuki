package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/crypto/bcrypt"
)

func newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", indexHandler)

		r.Post("/register", registerHandler)

		r.Get("/check-auth", checkAuthHandler)

		r.Get("/add-test-teacher", addTestTeacherHandler)
		r.Get("/add-test-student", addTestStudentHandler)
	})

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	users, err := getAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v\n", users)
	fmt.Fprint(w, "done\n")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var J struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&J)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := User{
		Name:         J.Name,
		Phone:        J.Phone,
		PasswordHash: hashPassword(J.Password),
	}

	id, err := addUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := createJWTtoken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, token)
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkAuthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := getJWTtokenFromCookies(r.Cookies())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "true")
}

func addTestTeacherHandler(w http.ResponseWriter, r *http.Request) {
	if err := addTestTeacher(); err != nil {
		e := fmt.Sprintf("cannot add test teacher: %v", err)
		http.Error(w, e, http.StatusInternalServerError)
	}
	fmt.Fprint(w, "success")
}

func addTestStudentHandler(w http.ResponseWriter, r *http.Request) {
	if err := addTestStudent(); err != nil {
		e := fmt.Sprintf("cannot add test student: %v", err)
		http.Error(w, e, http.StatusInternalServerError)
	}
	fmt.Fprint(w, "success")
}
