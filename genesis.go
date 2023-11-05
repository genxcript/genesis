package genesis

import (
	"fmt"
	"net/http"

	"github.com/genxcript/genesis/router"
)

func StartApp(routes []router.Route) {
	router := router.LoadRoutes(routes)
	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8081", router)
}
