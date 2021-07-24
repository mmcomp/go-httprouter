package httprouter

import (
	"net/http"
	"strings"
)

type Router struct {
	handlers  map[string]http.Handler
	delegates map[string]http.Handler
}

var Default Router = Router{
	handlers:  make(map[string]http.Handler),
	delegates: make(map[string]http.Handler),
}

func (receiver Router) Register(handler http.Handler, path string, method string) {
	receiver.handlers[method+":"+path] = handler
}

func (receiver Router) DelegatePath(handler http.Handler, path string, method string) {
	receiver.delegates[method+":"+path] = handler
}

func (receiver Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	for key, theHandler := range receiver.delegates {
		if strings.HasPrefix(method+":"+path, key) {
			theHandler.ServeHTTP(w, r)
			return
		}
	}
	handler, ok := receiver.handlers[method+":"+path]
	if ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Url Not Found!"))
	}
}
