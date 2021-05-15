package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
		r.Post("/login", loginHandler)
		r.Post("/create-teacher-profile", createTeacherProfile)

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

	_, err = addUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := createJWTtoken(u.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, token)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var J struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&J)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := getUserByPhone(J.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !checkPasswordHash(J.Password, user.PasswordHash) {
		http.Error(w, "passwords mismatch", http.StatusInternalServerError)
		return
	}

	token, err := createJWTtoken(user.Phone)
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

func createTeacherProfile(w http.ResponseWriter, r *http.Request) {
	var J struct {
		Subjects []string `json:"subjects"`
		Desc     string   `json:"desc"`
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&J)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := getJWTtokenFromCookies(r.Cookies())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	phone := claims["phone"].(string)

	user, err := getUserByPhone(phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Не забыть обновить в базе это поле
	user.Subjects = J.Subjects

	p := Profile{
		User: *user,
		Desc: J.Desc,
	}

	if err := addProfile(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "success")
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
