package session

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test all methods in one Go ;)
func TestSession(t *testing.T) {

	// create dummy server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set
		err := Set(w, r, "name", "john")
		if err != nil {
			t.Fatal(err)
		}
		// Get
		val, err := Get(r, "name")
		if value, ok := val.(string); ok {
			if value != "john" {
				t.Error("get value error")
			}
		}
		// Destroy
		err = Destroy(w, r)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintln(w, "Hello, World!")
	}))

	// use dummy server and send dummy request
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// make sure destroy worked
	// res.Cookies()[0] is the original cookie
	// res.Cookies()[1] is the same cookie with modified max-age
	cookie := res.Cookies()[1]
	past := time.Date(1970, 01, 01, 0, 0, 1, 0, time.UTC)
	if cookie.Expires != past {
		t.Error("session destroy error")
	}
}
