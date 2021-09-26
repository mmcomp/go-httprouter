package httprouter

import (
	"errors"
	"fmt"
	"net/http"
	"path"
)

type Router struct {
	handlers  map[string]map[string]http.Handler
	delegates map[string]map[string]http.Handler
}

var Default Router = Router{
	handlers:  make(map[string]map[string]http.Handler),
	delegates: make(map[string]map[string]http.Handler),
}

func (receiver Router) Register(handler http.Handler, path string, method string) error {
	var found bool
	_, found = receiver.handlers[path]
	if !found {
		receiver.handlers[path] = make(map[string]http.Handler)
	}
	_, found = receiver.handlers[path][method]
	if found {
		return errors.New("this path and method are already registered")
	}
	receiver.handlers[path][method] = handler
	return nil
}

func (receiver Router) DelegatePath(handler http.Handler, path string, method string) error {
	var found bool
	_, found = receiver.delegates[path]
	if !found {
		receiver.delegates[path] = make(map[string]http.Handler)
	}
	_, found = receiver.delegates[path][method]
	if found {
		return errors.New("this path and method are already registered")
	}
	receiver.delegates[path][method] = handler
	return nil
}

func (receiver Router) handler(method string, selectedPath string) (http.Handler, int) {
	fmt.Println("Router Handler  path ", method, ": ", selectedPath)
	var found bool
	var handlers map[string]http.Handler
	var handler http.Handler
	selectedPath = path.Clean(selectedPath)
	handlers, found = receiver.handlers[selectedPath]
	if found {
		handler, found = handlers[method]
		if found {
			return handler, 0
		}
		return nil, 405
	}
	fmt.Println("Delegates :")
	for p := range receiver.delegates {
		fmt.Println(p, " ", receiver.delegates[p])
	}
	handlers, found = receiver.delegates[selectedPath]
	fmt.Println("Delegate Handlers : ", found)
	if !found {
		var s string = selectedPath
		for {

			s = path.Clean(s)
			handlers, found = receiver.delegates[s]
			if found {
				break
			}
			handlers, found = receiver.delegates[s+"/"]
			if found {
				break
			}
			s = path.Dir(s)
			if s != "/" {
				s += "/"
			}

			if s == "/" {
				break
			}
		}
	}
	if found {
		handler, found = handlers[method]
		fmt.Printf("Delegate Handler : %t (%T) %#v", found, handler, handler)
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
