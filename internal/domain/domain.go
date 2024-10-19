package domain

import (
	"context"
	"os"

	"github.com/1ef7yy/go-kafka-poc/internal/storage"
	"github.com/1ef7yy/go-kafka-poc/pkg/logger"
)

type Domain interface {
}

type domain struct {
	log logger.Logger
	pg  storage.Postgres
}

func NewDomain(log logger.Logger) Domain {
	return &domain{
		log: log,
		pg:  *storage.NewPostgres(context.Background(), os.Getenv("POSTGRES_DSN"), log),
	}
}
