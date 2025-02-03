package view

import (
	"net/http"

	"github.com/1ef7yy/go-kafka-poc/pkg/logger"

	"github.com/1ef7yy/go-kafka-poc/internal/domain"
)

type View interface {
	ValidatePhone(w http.ResponseWriter, r *http.Request)
}

type view struct {
	log    logger.Logger
	domain domain.Domain
}

func NewView() View {
	log := logger.NewLogger(nil)
	return &view{
		log:    log,
		domain: domain.NewDomain(log),
	}
}
