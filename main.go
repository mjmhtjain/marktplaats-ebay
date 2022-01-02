package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mjmhtjain/marktplaats-ebay/src/router"
	"github.com/mjmhtjain/marktplaats-ebay/src/setup"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	setup.InitMongoClient()

	r := router.MuxRouter()
	http.Handle("/", r)

	log.Printf("Listing for requests at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
