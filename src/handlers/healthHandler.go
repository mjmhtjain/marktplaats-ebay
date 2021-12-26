package handlers

import "net/http"

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Healthy\n"))
}
