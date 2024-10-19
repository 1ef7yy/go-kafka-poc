package view

import (
	"github.com/1ef7yy/go-kafka-poc/internal/domain"
	"github.com/1ef7yy/go-kafka-poc/pkg/logger"
)

type view struct {
	log    logger.Logger
	domain domain.Domain
}

type View interface {
}

func NewView(log logger.Logger) View {
	return &view{
		log:    log,
		domain: domain.NewDomain(log),
	}
}
