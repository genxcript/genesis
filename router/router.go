package router

import "net/http"

type Route struct {
	Path        string
	Handler     func(w http.ResponseWriter, r *http.Request)
	Method      string
	Middlewares []func(w http.ResponseWriter, r *http.Request)
}
type Router struct {
	routes []Route
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if r.URL.Path == route.Path && r.Method == route.Method {
			for _, middleware := range route.Middlewares {
				middleware(w, r)
			}
			route.Handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func LoadRoutes(routes []Route) Router {
	return Router{
		routes: routes,
	}
}
