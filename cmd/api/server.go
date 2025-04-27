package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"

	"github.com/joho/godotenv"
)

type Form struct {
	Name []string
}

func main() {
	errr := godotenv.Load()
	if errr != nil {
		return
	}
	_, err := sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "")
		return
	}
	port := os.Getenv("API_PORT")

	cert := "../../cert.pem"
	key := "../../key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	// router:=router.Router()
	router := router.MainRouter()
	// secureMux := mw.SecurityHeaders(router)
	jwtMiddleware := mw.MiddlewareExcludePaths(mw.JWTMiddleware, "/execs/login")
	secureMux := jwtMiddleware(mw.SecurityHeaders(router))
	//Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
