package genesis

import (
	"fmt"
	"net/http"

	"github.com/genxcript/genesis/router"
)

func StartApp(routes []router.Route, host string) {
	router := router.LoadRoutes(routes)
	fmt.Printf("Starting server on %v\n", host)
	http.ListenAndServe(host, router)
}
