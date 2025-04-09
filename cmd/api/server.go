package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"restapi/internal/api/middlewares"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Printf("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("The ID is:", userID)

		fmt.Println("Query params", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("Sortby: %v, Sort order: %v, Key: %v", sortby, key, sortorder)

		w.Write([]byte("Hello MethodGet on Teachers Route"))
		// fmt.Println("Hello Get MethodGet on Teachers Route")
	case http.MethodPost:
		w.Write([]byte("Hello MethodPost  on Teachers Route"))
		fmt.Println("Hello Get MethodPost on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello MethodPut on Teachers Route"))
		fmt.Println("Hello Get MethodPut on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello MethodPatch on Teachers Route"))
		fmt.Println("Hello Get MethodPatch on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello MethodDelete on Teachers Route"))
		fmt.Println("Hello Get MethodDelete on Teachers Route")
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Students Route"))
	fmt.Printf("Hello Students Route")

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Execs Route"))
	fmt.Printf("Hello Execs Route")

}

func main() {
	port := ":3000"

	cert := "../../cert.pem"
	key := "../../key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
