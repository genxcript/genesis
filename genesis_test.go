package genesis

import (
	"net/http"
	"testing"
	"time"

	"github.com/genxcript/genesis/router"
)

func TestStartApp(t *testing.T) {
	routes := []router.Route{
		{
			Path:   "/test",
			Method: "GET",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("test passed"))
			}),
		},
	}

	go StartApp(routes, ":8083") // Start the app in a goroutine

	// Give the server a moment to start
	time.Sleep(time.Second)

	// Send a request to the server
	resp, err := http.Get("http://localhost:8083/test")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}
}
