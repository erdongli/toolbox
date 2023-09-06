package ip

import (
	"net/http"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/plain")

	ip := strings.TrimSpace(strings.Split(r.Header.Get("x-forwarded-for"), ",")[0])
	w.Write([]byte(ip))
}
