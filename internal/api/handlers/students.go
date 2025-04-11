package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Students Route"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello MethodGet on Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello MethodPost  on Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello MethodPut on Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello MethodPatch on Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello MethodDelete on Students Route"))
	}
}
