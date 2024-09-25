package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpMethod string

func (hm HttpMethod) Validate() error {
	switch hm {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return nil
	default:
		return errors.New("invalid http_helper method")
	}
}

type RouteDef struct {
	HttpMethod HttpMethod
	Pattern    string
	Handler    func(http.ResponseWriter, *http.Request)
}

func NewRouteDef(method HttpMethod, pattern string, handler func(http.ResponseWriter, *http.Request)) RouteDef {
	return RouteDef{
		HttpMethod: method,
		Pattern:    pattern,
		Handler:    handler,
	}

}

func LoadGroupOfRoutes(routeDefs []RouteDef) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	for _, def := range routeDefs {
		pattern := fmt.Sprintf("/%s", def.Pattern)
		switch def.HttpMethod {
		case http.MethodGet:
			router.Get(pattern, def.Handler)
		case http.MethodPost:
			router.Post(pattern, def.Handler)
		case http.MethodPut:
			router.Put(pattern, def.Handler)
		case http.MethodPatch:
			router.Patch(pattern, def.Handler)
		case http.MethodDelete:
			router.Delete(pattern, def.Handler)
		}
	}

	return router
}
