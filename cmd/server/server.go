package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type apiServer struct{}

func NewServer() *apiServer {
	return &apiServer{}
}

func (app *apiServer) Run() {
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
