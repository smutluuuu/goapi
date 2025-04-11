package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Execs Route"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello MethodGet on Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello MethodPost  on Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello MethodPut on Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello MethodPatch on Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello MethodDelete on Execs Route"))
	}

}
