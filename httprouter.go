package httprouter

import (
	"net/http"
)

type Router struct {
	handlers  map[string]map[string]http.Handler
	delegates map[string]map[string]http.Handler
}

var Default Router = Router{
	handlers:  make(map[string]map[string]http.Handler),
	delegates: make(map[string]map[string]http.Handler),
}

func (receiver Router) Register(handler http.Handler, path string, method string) {
	_, found := receiver.handlers[path]
	if !found {
		receiver.handlers[path] = make(map[string]http.Handler)
	}
	receiver.handlers[path][method] = handler
}

func (receiver Router) DelegatePath(handler http.Handler, path string, method string) {
	_, found := receiver.delegates[path]
	if !found {
		receiver.delegates[path] = make(map[string]http.Handler)
	}
	receiver.delegates[path][method] = handler
}

func (receiver Router) handler(method string, path string) (http.Handler, int) {
	var found bool
	var handlers map[string]http.Handler
	var handler http.Handler
	handlers, found = receiver.handlers[path]
	if found {
		handler, found = handlers[method]
		if found {
			return handler, 0
		}
		return nil, 405
	}

	handlers, found = receiver.delegates[path]
	if !found {
		handlers, found = receiver.delegates[path+"/"]
	}
	if found {
		handler, found = handlers[method]
		if found {
			return handler, 0
		}
		return nil, 405
	}

	return nil, 404
}

func (receiver Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	handler, status := receiver.handler(method, path)

	if nil == handler {
		switch status {
		case 405:
			w.WriteHeader(405)
			w.Write([]byte("405 Method Not Allowed"))
			return
		default:
			w.WriteHeader(404)
			w.Write([]byte("404 Not Found!"))
			return
		}
	}

	handler.ServeHTTP(w, r)
}
