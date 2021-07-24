package httprouter

import (
	"fmt"
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
	fmt.Println("Checking for delegate path : ", method+":"+path)
	for key, theHandler := range receiver.delegates {
		fmt.Println("Compare ", key)
		if strings.HasPrefix(method+":"+path, key) {
			fmt.Println("OK!")
			theHandler.ServeHTTP(w, r)
			return
		}
	}
	fmt.Println("Checking for path : ", method+":"+path)
	handler, ok := receiver.handlers[method+":"+path]
	if ok {
		fmt.Println("OK!")
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Url Not Found!"))
	}
}
