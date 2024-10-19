package v1

import (
	"net/http"
)

func (v *Router) ApiEndpoint() http.Handler {

	apimux := http.NewServeMux()

	apimux.Handle("GET /some_endpoint", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	return http.StripPrefix("/api", apimux)

}
