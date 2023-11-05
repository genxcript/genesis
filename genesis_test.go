package genesis

import (
	"net/http"
	"net/http/httptest"
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

	ConfigSet("app", "host", ":8083")
	go StartApp(routes) // Start the app in a goroutine

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

func TestView(t *testing.T) {
	ConfigSet("app", "views_dir", "tests/stubs")
	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Call the View function, passing in the ResponseRecorder as the http.ResponseWriter.
	View(rr, "view", nil)

	// Now we can check the recorded response!

	// Check the response status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `content of tests/stubs/view.html`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestConfig(t *testing.T) {

	configObject.Set("testPath", "testKey", "testValue")

	// Call the Config function
	value := Config("testPath", "testKey")

	// Check the returned value
	if value != "testValue" {
		t.Errorf("Config returned wrong value: got %v want %v", value, "testValue")
	}
}

func TestConfigSet(t *testing.T) {

	// Call the ConfigSet function
	ConfigSet("testPath", "testKey", "testValue")

	// Check the value was set correctly
	if Config("testPath", "testKey") != "testValue" {
		t.Errorf("ConfigSet did not set the correct value: got %v want %v", Config("testPath", "testKey"), "testValue")
	}
}
