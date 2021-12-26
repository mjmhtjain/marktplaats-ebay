package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}
