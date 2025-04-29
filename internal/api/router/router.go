package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := teachersRouter()
	sRouter := studentsRouter()
	eRouter := execsRouter()

	sRouter.Handle("/", eRouter)
	tRouter.Handle("/", sRouter)
	return tRouter
}
