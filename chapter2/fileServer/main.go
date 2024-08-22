package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.ServeFiles("/static/*filepath",
		http.Dir("/home/wendersom/Desktop/LIVROS"))

	log.Fatal(http.ListenAndServe(":8000", router))
}
