package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	r := router()
	http.Handle("/", r)

	log.Printf("Listing for requests at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)

	return r
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Hello, world!\n")
	w.Write([]byte("Hello, world!\n"))
}
