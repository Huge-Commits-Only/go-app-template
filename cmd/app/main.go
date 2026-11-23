// Command app is the starting point of a new Huge Commits Only project.
//
// Configuration (environment): PORT (default 8080).
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() { os.Exit(run(os.Args[1:], os.Getenv)) }

// run is the whole program, arguments and environment in, exit code out, so a test can drive it.
func run(args []string, getenv func(string) string) int {
	port := getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// `app -health` is the container health check: the image has no shell or curl to do it with.
	if len(args) > 0 && args[0] == "-health" {
		resp, err := http.Get("http://127.0.0.1:" + port + "/healthz")
		if err != nil {
			return 1
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusOK {
			return 1
		}
		return 0
	}
	log.Printf("app: http://0.0.0.0:%s/", port)
	srv := &http.Server{Addr: ":" + port, Handler: handler(), ReadHeaderTimeout: 10 * time.Second}
	log.Print(srv.ListenAndServe())
	return 1
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("Hello from Huge Commits Only"))
	})
	return mux
}
