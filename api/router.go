package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	teachers, err := getAllTeachers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v\n", teachers)

	students, err := getAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v\n", students)

	fmt.Fprint(w, "done\n")
}

func registerTeacherHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	password := strings.TrimSpace(r.FormValue("password"))

	t := Teacher{
		Name:         name,
		Phone:        phone,
		PasswordHash: hashPassword(password),
	}

	id, err := addTeacher(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := createJWTtoken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().AddDate(0, 1, 0),
		SameSite: 3,
	}
	http.SetCookie(w, &c)

	fmt.Fprint(w, "success")
}

func registerStudentHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	password := strings.TrimSpace(r.FormValue("password"))

	s := Student{
		Name:         name,
		Phone:        phone,
		PasswordHash: hashPassword(password),
	}

	id, err := addStudent(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := createJWTtoken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().AddDate(0, 1, 0),
		SameSite: 3,
	}
	http.SetCookie(w, &c)

	fmt.Fprint(w, "success")
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
