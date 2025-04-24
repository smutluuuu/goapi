package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func studentsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudentHandler)
	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudentHandler)

	return mux
}
