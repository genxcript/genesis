package genesis

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/genxcript/genesis/config"
	"github.com/genxcript/genesis/router"
)

var configObject = config.ConfigObject{}

func StartApp(routes []router.Route) {
	host := Config("app", "host")
	if host == nil {
		host = ":3000"
	}

	router := router.LoadRoutes(routes)
	fmt.Printf("Starting server on %v\n", host)
	http.ListenAndServe(host.(string), router)
}

func View(w http.ResponseWriter, view string, data interface{}) {
	viewPath := Config("app", "views_dir")
	if viewPath == nil {
		viewPath = "views"
	}

	template.Must(template.ParseFiles(fmt.Sprintf("%v/%v.html", viewPath, view))).Execute(w, data)
}

func Config(path string, key string) any {
	return configObject.Get(path, key)
}

func ConfigSet(path string, key string, value any) {
	configObject.Set(path, key, value)
}
