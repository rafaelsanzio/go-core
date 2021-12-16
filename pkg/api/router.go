package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelsanzio/go-core/pkg/api/handlers"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	for _, r := range routes {
		router.Methods(r.Methods...).Path(r.Path).Name(r.Name).HandlerFunc(r.Handler)
	}

	return router
}

type Route struct {
	Name    string
	Methods []string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	{
		Name:    "Health OK",
		Methods: []string{http.MethodGet},
		Path:    "/ok",
		Handler: handlers.HandleAdapter(handlers.HandleOK),
	},
	{
		Name:    "Post User",
		Methods: []string{http.MethodPost},
		Path:    "/users",
		Handler: handlers.HandleAdapter(handlers.HandlePostUser),
	},
}