package main

import (
	"net/http"
	"os"

	"github.com/1ef7yy/go-kafka-poc/internal/routes"
	"github.com/1ef7yy/go-kafka-poc/internal/view"
	"github.com/1ef7yy/go-kafka-poc/pkg/logger"
)

func main() {
	logger := logger.NewLogger(nil)

	logger.Info("starting server...")

	view := view.NewView()

	logger.Info("initializing router...")

	mux := routes.InitRouter(view)

	logger.Info("server started on " + os.Getenv("SERVER_ADDRESS"))
	if err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), mux); err != nil {
		logger.Fatal(err.Error())
	}
}
