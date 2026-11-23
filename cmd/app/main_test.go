package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()
	for path, want := range map[string]int{"/": 200, "/healthz": 200, "/missing": 404} {
		resp, err := http.Get(srv.URL + path)
		if err != nil {
			t.Fatal(err)
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != want {
			t.Errorf("GET %s: %d, want %d", path, resp.StatusCode, want)
		}
	}
}

// The container health check passes against a running server and fails without one.
func TestHealthCheck(t *testing.T) {
	srv := httptest.NewServer(handler())
	u, _ := url.Parse(srv.URL)
	env := func(port string) func(string) string { return func(string) string { return port } }
	if code := run([]string{"-health"}, env(u.Port())); code != 0 {
		t.Fatalf("healthy server: exit %d", code)
	}
	srv.Close()
	if code := run([]string{"-health"}, env(u.Port())); code != 1 {
		t.Fatalf("stopped server: exit %d", code)
	}
}

// A server that cannot start (an invalid port) ends with exit code 1 instead of hanging.
func TestServerFailsToStart(t *testing.T) {
	if code := run(nil, func(string) string { return "-1" }); code != 1 {
		t.Fatalf("exit %d", code)
	}
}
