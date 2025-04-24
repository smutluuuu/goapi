package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func execsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /execs", handlers.GetExecsHandler)
	mux.HandleFunc("POST /execs", handlers.AddExecsHandler)
	mux.HandleFunc("PATCH /execs", handlers.PatchExecsHandler)

	mux.HandleFunc("GET /execs/{id}", handlers.GetOneExecHandler)
	mux.HandleFunc("PATCH /execs/{id}", handlers.PatchOneExecHandler)
	mux.HandleFunc("DELETE /execs/{id}", handlers.DeleteOneExecHandler)
	mux.HandleFunc("POST /execs/{id}/updatepassword", handlers.GetExecsHandler)

	mux.HandleFunc("POST /execs/login", handlers.GetExecsHandler)
	mux.HandleFunc("POST /execs/logout", handlers.GetExecsHandler)
	mux.HandleFunc("POST /execs/forgotpassword", handlers.GetExecsHandler)
	mux.HandleFunc("POST /execs/resetpassword/reset/{resetcode}", handlers.GetExecsHandler)

	return mux

}
