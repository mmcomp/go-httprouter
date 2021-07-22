package httprouter

import (
	"net/http"
)

type Router struct {
	handlers map[string]http.Handler
}

var Default Router = Router{
	handlers: make(map[string]http.Handler),
}

func (receiver Router) Register(handler http.Handler, path string, method string) {
	receiver.handlers[method+":"+path] = handler
}

func (receiver Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	handler, ok := receiver.handlers[method+":"+path]
	if ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Url Not Found!"))
	}
}
