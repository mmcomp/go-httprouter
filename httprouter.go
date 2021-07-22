package httprouter

import (
	"net/http"
)

type Router struct {
}

func (receiver *Router) Register(handler http.Handler, path string, method string) {
	// ToDO
}

func (receiver *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ToDO
}
