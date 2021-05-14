package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register-teacher", registerTeacherHandler)
	http.HandleFunc("/register-student", registerStudentHandler)

	http.HandleFunc("/check-auth", checkAuthHandler)

	http.HandleFunc("/add-test-teacher", addTestTeacherHandler)
	http.HandleFunc("/add-test-student", addTestStudentHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
