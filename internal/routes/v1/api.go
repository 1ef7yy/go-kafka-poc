package v1

import (
	"net/http"
)

func (v *Router) ApiEndpoint() http.Handler {

	apimux := http.NewServeMux()

	apimux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	return http.StripPrefix("/api", apimux)

}
