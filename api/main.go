package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/", indexHandler)
	http.HandleFunc("/api/v1/register-teacher", registerTeacherHandler)
	http.HandleFunc("/api/v1/register-student", registerStudentHandler)

	http.HandleFunc("/api/v1/check-auth", checkAuthHandler)

	http.HandleFunc("/api/v1/add-test-teacher", addTestTeacherHandler)
	http.HandleFunc("/api/v1/add-test-student", addTestStudentHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
