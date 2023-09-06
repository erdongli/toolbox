package server

import (
	"net/http"

	"github.com/erdongli/toolbox/internal/ip"
)

func Serve() error {
	http.HandleFunc("/ip", ip.Handle)

	return http.ListenAndServe(":8080", nil)
}
