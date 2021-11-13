package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shubham-umate/mongoapi/router"
)

func main() {
	fmt.Println("Lets learn mongo api")
	r := router.Router()

	fmt.Println("Server is getting started")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening on port 4000")
}
