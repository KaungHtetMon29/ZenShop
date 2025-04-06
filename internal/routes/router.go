package routes

import "net/http"
type router struct { 
	routes map[string]map[string]http.HandlerFunc
}
func NewRouter() *router {
	return &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}
func InitializeRoutes() *router{
	r:=NewRouter()
	r.AddRoute("GET", "/users", nil)
	return r
}
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path:=req.URL.Path
	method:=req.Method
	if handlers, ok := r.routes[path]; ok {
		if handler, ok := handlers[method]; ok {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *router) AddRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path]==nil{
		r.routes[path]=make(map[string]http.HandlerFunc)
	}
	r.routes[path][method]=handler
}