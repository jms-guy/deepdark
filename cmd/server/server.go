package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/jms-guy/deepdark/cmd/server/internal/logger"
)

type AppServer struct {
	Logger logger.Logger
}

func NewServer() *AppServer {
	return &AppServer{}
}

// Main server setup and run function
func (app *AppServer) Run() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	mux.HandleFunc("GET /health", app.healthHandle)

	fmt.Println("Listening at ", server.Addr)
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
