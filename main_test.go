package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type redirTest struct {
	method  string
	url     string
	dest    string
	rstatus int
	rstring string
}

func TestHealthCheckHandler(t *testing.T) {
	var redirTests = []redirTest{
		{"GET", "http://foo/", "443", 301, "https://foo/"},
		{"GET", "http://foo/", "8080", 301, "https://foo:8080/"},
		{"GET", "http://foo/bar", "8080", 301, "https://foo:8080/bar"},
		{"GET", "http://foo/bar?baz=true", "8080", 301, "https://foo:8080/bar?baz=true"},
		{"POST", "http://foo/bar", "443", 308, ""},
	}

	for _, tt := range redirTests {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(tt.method, tt.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := RedirectHandler{Destination: tt.dest}

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tt.rstatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tt.rstatus)
		}

		// Check the response body is what we expect.
		expected := ""
		if tt.rstring != "" {
			expected = fmt.Sprintf("<a href=\"%s\">Moved Permanently</a>.\n\n", tt.rstring)
		}
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

func TestParseOptions(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"redirect", "-port=81", "-destination=8443"}
	expected := Options{
		ShowHelp:    false,
		Source:      "81",
		Destination: "8443",
	}
	actual := parseOptions()

	if actual != expected {
		t.Errorf("Test failed, expected: '{%t,%s,%s}', got:  '{%t,%s,%s}'",
			expected.ShowHelp, expected.Source, expected.Destination,
			actual.ShowHelp, actual.Source, actual.Destination)
	}
}
