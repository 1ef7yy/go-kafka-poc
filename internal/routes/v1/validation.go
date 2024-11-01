package v1

import (
	"net/http"
)

func (v *Router) Validation() http.Handler {

	apimux := http.NewServeMux()

	apimux.Handle("GET /validatePhone", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Example: http://host_url/api/validation/validatePhone?phone_number=+79527073822

		v.View.ValidatePhone(w, r)
	}))

	return http.StripPrefix("/api/validation", apimux)
}
