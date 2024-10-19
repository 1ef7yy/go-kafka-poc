package v1

import (
	"net/http"

	"github.com/1ef7yy/go-kafka-poc/internal/view"
)

type Router struct {
	View view.View
}

func NewRouter(view view.View) *Router {
	return &Router{
		View: view,
	}
}

func (v *Router) Api() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	return http.StripPrefix("/api", mux)
}
