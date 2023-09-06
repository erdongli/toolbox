package server

import (
	"net/http"

	"github.com/erdongli/toolbox/internal/ip"
)

func Serve() error {
	// Liveness and readiness probes
	http.HandleFunc("/live", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/ready", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ip", ip.Handle)

	return http.ListenAndServe(":8080", nil)
}
