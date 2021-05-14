package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/add-test-teacher", func(w http.ResponseWriter, r *http.Request) {
		if err := addTestTeacher(); err != nil {
			e := fmt.Sprintf("cannot add test teacher: %v", err)
			http.Error(w, e, http.StatusInternalServerError)
		}
		fmt.Fprint(w, "success\n")
	})

	http.HandleFunc("/add-test-student", func(w http.ResponseWriter, r *http.Request) {
		if err := addTestStudent(); err != nil {
			e := fmt.Sprintf("cannot add test student: %v", err)
			http.Error(w, e, http.StatusInternalServerError)
		}
		fmt.Fprint(w, "success\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
