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
	completePath := method+":"+path
	fmt.Println("HttpRouter : ", completePath)
	handler, ok := receiver.handlers[completePath]
	if ok {
		fmt.Println("Found Strict Handler")
		handler.ServeHTTP(w, r)
		return
	}

	for key, theHandler := range receiver.delegates {
		fmt.Println("Delegate Checking ", key+"/")
		if strings.HasPrefix(completePath, key+"/") {
			fmt.Println("Found Delegate Handler")
			theHandler.ServeHTTP(w, r)
			return
		}
		fmt.Println("Delegate Checking ", key)
		if strings.HasPrefix(completePath, key) {
			fmt.Println("Found Delegate Handler")
			theHandler.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Url Not Found!"))
}
