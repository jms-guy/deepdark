package main

import (
	"net/http"
)

func (app *AppServer) healthHandle(w http.ResponseWriter, r *http.Request) {
	app.respondWithJSON(w, http.StatusOK, "OK")
}
